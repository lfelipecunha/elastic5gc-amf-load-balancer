package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	CauseTransportPresentTransportResourceUnavailable aper.Enumerated = 0
	CauseTransportPresentUnspecified                  aper.Enumerated = 1
)

type CauseTransport struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
