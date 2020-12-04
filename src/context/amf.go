package context

import (
	"amfLoadBalancer/lib/ngap/ngapSctp"
	"amfLoadBalancer/src/logger"
	"fmt"
	"net"

	"github.com/ishidawataru/sctp"
)

type AmfData struct {
	IP   string
	Port int
	ID   string
}

type Amf struct {
	/* socket Connect*/
	Conn    net.Conn
	AmfData *AmfData
}

func NewAmf(amfData *AmfData, ranIP string, ranPort int) (*Amf, error) {
	var amf Amf
	var err error
	amf.AmfData = amfData
	amf.Conn, err = ConnectToAmf(amfData.IP, ranIP, amfData.Port, ranPort)
	if err != nil {
		logger.ContextLog.Error("Unable to connect to AMF")
		return nil, err
	}

	return &amf, nil
}

func getNgapIp(amfIP, ranIP string, amfPort, ranPort int) (amfAddr, ranAddr *sctp.SCTPAddr, err error) {
	ips := []net.IPAddr{}
	// se der um erro != nill entra no if.
	if ip, err1 := net.ResolveIPAddr("ip", amfIP); err1 != nil {
		err = fmt.Errorf("Error resolving address '%s': %v", amfIP, err1)
		return
	} else {
		ips = append(ips, *ip)
	}
	amfAddr = &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    amfPort,
	}
	ips = []net.IPAddr{}
	if ip, err1 := net.ResolveIPAddr("ip", ranIP); err1 != nil {
		err = fmt.Errorf("Error resolving address '%s': %v", ranIP, err1)
		return
	} else {
		ips = append(ips, *ip)
	}
	ranAddr = &sctp.SCTPAddr{
		IPAddrs: ips,
		Port:    ranPort,
	}
	return
}

func ConnectToAmf(amfIP, ranIP string, amfPort, ranPort int) (*sctp.SCTPConn, error) {
	amfAddr, ranAddr, err := getNgapIp(amfIP, ranIP, amfPort, ranPort)
	if err != nil {
		return nil, err
	}
	conn, err := sctp.DialSCTP("sctp", ranAddr, amfAddr)
	if err != nil {
		return nil, err
	}
	info, _ := conn.GetDefaultSentParam()
	info.PPID = ngapSctp.NGAP_PPID
	err = conn.SetDefaultSentParam(info)
	if err != nil {
		return nil, err
	}
	return conn, nil
}
