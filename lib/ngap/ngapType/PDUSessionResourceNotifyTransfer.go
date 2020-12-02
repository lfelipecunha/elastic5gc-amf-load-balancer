package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceNotifyTransfer struct {
	QosFlowNotifyList   *QosFlowNotifyList                                                `aper:"optional"`
	QosFlowReleasedList *QosFlowList                                                      `aper:"optional"`
	IEExtensions        *ProtocolExtensionContainerPDUSessionResourceNotifyTransferExtIEs `aper:"optional"`
}
