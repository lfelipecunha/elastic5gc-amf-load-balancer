package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type AMFUENGAPID struct {
	Value int64 `aper:"valueLB:0,valueUB:1099511627775"`
}
