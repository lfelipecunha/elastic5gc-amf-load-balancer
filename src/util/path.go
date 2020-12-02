//+build !debug

package util

import (
	"amfLoadBalancer/lib/path_util"
)

var AmfLogPath = path_util.GoamfLoadBalancerPath("amfLoadBalancer/amfsslkey.log")
var AmfPemPath = path_util.GoamfLoadBalancerPath("amfLoadBalancer/support/TLS/amf.pem")
var AmfKeyPath = path_util.GoamfLoadBalancerPath("amfLoadBalancer/support/TLS/amf.key")
var DefaultAmfConfigPath = path_util.GoamfLoadBalancerPath("amfLoadBalancer/config/amfcfg.conf")
