package context

import (
	"amfLoadBalancer/lib/ngap/ngapSctp"
	"amfLoadBalancer/src/logger"
	"fmt"
	"net"

	"github.com/ishidawataru/sctp"
)

var RanPort = 9847

type Amf struct {
	/* socket Connect*/
	Conn net.Conn
}

func NewAmf(amfIP string, amfPort int, ranIP string) *Amf {
	var amf Amf
	var err error

	amf.Conn, err = ConnectToAmf(amfIP, ranIP, amfPort, RanPort)
	if err != nil {
		logger.ContextLog.Error("Unable to connect to AMF")
	}
	RanPort++

	return &amf
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
