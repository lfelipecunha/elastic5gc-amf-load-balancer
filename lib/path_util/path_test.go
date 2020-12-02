package path_util

import (
	"amfLoadBalancer/lib/path_util/logger"
	"testing"
)

func TestFree5gcPath(t *testing.T) {
	logger.PathLog.Infoln(GoamfLoadBalancerPath("amfLoadBalancer/abcdef/abcdef.pem"))
}
