package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type NumberOfBroadcastsRequested struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}
