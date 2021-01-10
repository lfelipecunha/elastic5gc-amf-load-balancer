package service

import (
	"amfLoadBalancer/src/consumer"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
)

func UpdateAmfList() {
	var controlAmfs []*context.AmfData
	result, _ := consumer.GetAMFInstances(context.AMF_Self().NrfUri)
	balancer := context.AMF_Self().Balancer
	amfs := balancer.GetAmfs()

	controlAmfs = amfs
	logger.AppLog.Debugf("Founded %d AMFs", len(result.NfInstances))
	for i := 0; i < len(result.NfInstances); i++ {
		ip := result.NfInstances[i].Ipv4Addresses[0]
		id := result.NfInstances[i].NfInstanceId
		logger.AppLog.Debugf("AMF[%s] = %s", id, ip)
		add := true
		for index, existentAmf := range amfs {
			if id == existentAmf.ID {
				add = false
				controlAmfs = append(controlAmfs[:index], controlAmfs[index+1:]...)
				break
			}
		}
		if add {
			balancer.AddAmf(&context.AmfData{IP: ip, Port: 38412, ID: id, Blocked: false})
		}
	}

	for _, amf := range controlAmfs {
		balancer.RemoveAmf(amf.ID)
	}
}

func BlockAmf(url string) {
	balancer := context.AMF_Self().Balancer
	balancer.BlockAmf(url)

}
