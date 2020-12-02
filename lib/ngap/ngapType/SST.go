package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type SST struct {
	Value aper.OctetString `aper:"sizeLB:1,sizeUB:1"`
}
