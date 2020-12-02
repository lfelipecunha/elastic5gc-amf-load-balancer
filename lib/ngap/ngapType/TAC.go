package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TAC struct {
	Value aper.OctetString `aper:"sizeLB:3,sizeUB:3"`
}
