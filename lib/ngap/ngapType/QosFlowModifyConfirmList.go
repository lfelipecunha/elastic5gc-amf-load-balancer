package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct QosFlowModifyConfirmList */
/* QosFlowModifyConfirmItem */
type QosFlowModifyConfirmList struct {
	List []QosFlowModifyConfirmItem `aper:"valueExt,sizeLB:1,sizeUB:64"`
}
