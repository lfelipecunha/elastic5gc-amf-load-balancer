package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type ERABID struct {
	Value int64 `aper:"valueExt,valueLB:0,valueUB:15"`
}
