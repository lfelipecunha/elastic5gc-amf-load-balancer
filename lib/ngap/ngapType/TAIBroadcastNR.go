package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct TAIBroadcastNR */
/* TAIBroadcastNRItem */
type TAIBroadcastNR struct {
	List []TAIBroadcastNRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
