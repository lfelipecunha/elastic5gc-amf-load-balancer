package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowInformationList */
/* QosFlowInformationItem */
type QosFlowInformationList struct {
	List []QosFlowInformationItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
