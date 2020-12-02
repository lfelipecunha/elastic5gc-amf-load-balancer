package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type LocationReportingReferenceID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:64"`
}
