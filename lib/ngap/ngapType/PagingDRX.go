package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	PagingDRXPresentV32  aper.Enumerated = 0
	PagingDRXPresentV64  aper.Enumerated = 1
	PagingDRXPresentV128 aper.Enumerated = 2
	PagingDRXPresentV256 aper.Enumerated = 3
)

type PagingDRX struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:3"`
}
