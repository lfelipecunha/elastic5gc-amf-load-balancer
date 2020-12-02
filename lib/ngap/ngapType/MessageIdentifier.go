package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type MessageIdentifier struct {
	Value aper.BitString `aper:"sizeLB:16,sizeUB:16"`
}
