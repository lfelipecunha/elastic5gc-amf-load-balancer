package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct CellIDBroadcastNR */
/* CellIDBroadcastNRItem */
type CellIDBroadcastNR struct {
	List []CellIDBroadcastNRItem `aper:"valueExt,sizeLB:1,sizeUB:65535"`
}
