package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	DataForwardingNotPossiblePresentDataForwardingNotPossible aper.Enumerated = 0
)

type DataForwardingNotPossible struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
