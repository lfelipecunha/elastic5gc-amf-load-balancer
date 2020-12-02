package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type AMFRegionID struct {
	Value aper.BitString `aper:"sizeLB:8,sizeUB:8"`
}
