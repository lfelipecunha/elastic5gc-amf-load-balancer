package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TransportLayerAddress struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:1,sizeUB:160"`
}
