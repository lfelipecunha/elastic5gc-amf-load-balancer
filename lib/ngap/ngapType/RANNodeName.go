package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type RANNodeName struct {
	Value string `aper:"sizeExt,sizeLB:1,sizeUB:150"`
}
