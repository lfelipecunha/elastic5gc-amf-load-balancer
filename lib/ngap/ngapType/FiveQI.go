package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type FiveQI struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:255"`
}
