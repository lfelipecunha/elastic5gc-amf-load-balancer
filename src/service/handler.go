package service

import (
	"amfLoadBalancer/src/consumer"
	"amfLoadBalancer/src/logger"
	"encoding/hex"
	"io"
	"net"
	"net/http"
	"strconv"
	"sync"

	"github.com/ishidawataru/sctp"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
)

type Handler func(conn net.Conn, msg []byte)

const NGAP_PPID uint32 = 0x3c000000
const readBufSize uint32 = 8192

var sctpListener *sctp.SCTPListener
var connections sync.Map

func RunSCtp(addresses []string, port int, msgHandler Handler) {
	ips := []net.IPAddr{}

	for _, addr := range addresses {
		if netAddr, err := net.ResolveIPAddr("ip", addr); err != nil {
			logger.NgapLog.Errorf("Error resolving address '%s': %v\n", addr, err)
		} else {
			logger.NgapLog.Debugf("Resolved address '%s' to %s\n", addr, netAddr)
			ips = append(ips, *netAddr)
		}
	}

	addr := &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    port,
	}

	listenAndServeSctp(addr, msgHandler)
}

func RunHttp(address string, port int, nfuri string) {
	address = address + ":" + strconv.Itoa(port)
	listenAndServeHttp(address)
	consumer.SubscribeAmfsChanges(nfuri, "http://"+address)
}

func listenAndServeSctp(addr *sctp.SCTPAddr, msgHandler Handler) {
	initMsg := sctp.InitMsg{NumOstreams: 3, MaxInstreams: 5, MaxAttempts: 4, MaxInitTimeout: 8}

	if listener, err := sctp.ListenSCTPExt("sctp", addr, initMsg); err != nil {
		logger.NgapLog.Errorf("Failed to listen: %+v", err)
		return
	} else {
		sctpListener = listener
	}

	logger.NgapLog.Infof("Listen on %s", sctpListener.Addr())

	for {
		var conn *sctp.SCTPConn
		if newConn, err := sctpListener.AcceptSCTP(); err != nil {
			logger.NgapLog.Errorf("Failed to accept: %+v", err)
			continue
		} else {
			conn = newConn
		}

		var info *sctp.SndRcvInfo
		if infoTmp, err := conn.GetDefaultSentParam(); err != nil {
			logger.NgapLog.Errorf("Failed to accept: %+v", err)
			if err = conn.Close(); err != nil {
				logger.NgapLog.Errorf("Close error: %+v", err)
			}
		} else {
			info = infoTmp
		}

		info.PPID = NGAP_PPID
		if err := conn.SetDefaultSentParam(info); err != nil {
			logger.NgapLog.Errorf("Failed to accept: %+v", err)
			if err = conn.Close(); err != nil {
				logger.NgapLog.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			logger.NgapLog.Debugf("Set default sent param[value: %+v] successfully", info)
		}

		if err := conn.SubscribeEvents(sctp.SCTP_EVENT_DATA_IO); err != nil {
			logger.NgapLog.Errorf("Failed to accept: %+v", err)
			if err = conn.Close(); err != nil {
				logger.NgapLog.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			logger.NgapLog.Debugln("Subscribe SCTP event DATA_IO successfully")
		}

		if err := conn.SetReadBuffer(int(readBufSize)); err != nil {
			logger.NgapLog.Errorf("Failed to accept: %+v", err)
			if err = conn.Close(); err != nil {
				logger.NgapLog.Errorf("Close error: %+v", err)
			}
			continue
		} else {
			logger.NgapLog.Debugf("Set read buffer to %d bytes", readBufSize)
		}
		remoteAddr := conn.RemoteAddr()
		if remoteAddr == nil {
			logger.NgapLog.Errorf("Connection closed!")
			continue
		}
		logger.NgapLog.Infof("[AMF] SCTP Accept from: %s", remoteAddr.String())

		connections.Store(conn, conn)
		go func() {
			if err := handleConnection(conn, readBufSize, msgHandler); err != nil {
				logger.NgapLog.Errorf("Handle connection[addr: %+v] error: %+v", conn.RemoteAddr(), err)
			}
			// if AMF call Stop(), then conn.Close() will return "bad file descriptor" error
			// because conn has been closed inside Stop()
			if err := conn.Close(); err != nil && err.Error() != "bad file descriptor" {
				logger.NgapLog.Errorf("close connection error: %+v", err)
			}
			connections.Delete(conn)
		}()
	}
}

func Stop() {
	logger.NgapLog.Infof("Close SCTP server...")
	if err := sctpListener.Close(); err != nil {
		logger.NgapLog.Error(err)
		logger.NgapLog.Infof("SCTP server may not close normally.")
	}

	connections.Range(func(key, value interface{}) bool {
		conn := value.(net.Conn)
		if err := conn.Close(); err != nil {
			logger.NgapLog.Error(err)
		}
		return true
	})

	logger.NgapLog.Infof("SCTP server closed")
}

func handleConnection(conn *sctp.SCTPConn, bufsize uint32, msgHandler Handler) error {
	for {
		buf := make([]byte, bufsize)

		n, info, err := conn.SCTPRead(buf)
		if err != nil {
			if err == io.EOF || err == io.ErrUnexpectedEOF {
				logger.NgapLog.Debugln("Read EOF from client")
				return nil
			} else {
				return err
			}
		}

		if info == nil || info.PPID != NGAP_PPID {
			logger.NgapLog.Warnln("Received SCTP PPID != 60, discard this packet")
			continue
		}

		logger.NgapLog.Tracef("Read %d bytes", n)
		logger.NgapLog.Tracef("Packet content:\n%+v", hex.Dump(buf[:n]))

		// TODO: concurrent on per-UE message
		msgHandler(conn, buf[:n])
	}
}

func listenAndServeHttp(uri string) {
	// Create a server on port 8000
	h2s := &http2.Server{}
	handler := http.HandlerFunc(handleHttp)

	srv := &http.Server{Addr: uri, Handler: h2c.NewHandler(handler, h2s)}

	// Exactly how you would run an HTTP/1.1 server with TLS connection.
	logger.AppLog.Info("Serving on " + uri)
	go srv.ListenAndServe()
}

func handleHttp(w http.ResponseWriter, r *http.Request) {
	// Log the request protocol
	logger.HttpLog.Infof("Got connection: %s", r.Proto)
	UpdateAmfList()
	w.WriteHeader(http.StatusNoContent)
	// Send a message back to the client
	w.Write([]byte(""))
}
