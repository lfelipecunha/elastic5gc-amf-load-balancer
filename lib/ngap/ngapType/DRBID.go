package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type DRBID struct {
	Value int64 `aper:"valueExt,valueLB:1,valueUB:32"`
}
