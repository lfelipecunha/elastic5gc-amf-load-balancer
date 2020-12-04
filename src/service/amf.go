package service

import (
	"amfLoadBalancer/src/consumer"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
)

func UpdateAmfList() {
	var i int
	result, _ := consumer.GetAMFInstances(context.AMF_Self().NrfUri)
	logger.AppLog.Debugf("Found %i AMFs", len(result.NfInstances))
	for i = 0; i < len(result.NfInstances); i++ {
		ip := result.NfInstances[i].Ipv4Addresses[0]
		id := result.NfInstances[i].NfInstanceId
		logger.AppLog.Debugf("AMF[%s] = %s", id, ip)

		context.AMF_Self().Balancer.AddAmf(&context.AmfData{IP: ip, Port: 38412, ID: id})
	}
}
