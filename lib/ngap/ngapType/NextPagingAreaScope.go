package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	NextPagingAreaScopePresentSame    aper.Enumerated = 0
	NextPagingAreaScopePresentChanged aper.Enumerated = 1
)

type NextPagingAreaScope struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
