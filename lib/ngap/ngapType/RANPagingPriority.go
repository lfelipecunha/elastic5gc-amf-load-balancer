package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type RANPagingPriority struct {
	Value int64 `aper:"valueLB:1,valueUB:256"`
}
