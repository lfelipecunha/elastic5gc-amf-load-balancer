package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type WarningType struct {
	Value aper.OctetString `aper:"sizeLB:2,sizeUB:2"`
}
