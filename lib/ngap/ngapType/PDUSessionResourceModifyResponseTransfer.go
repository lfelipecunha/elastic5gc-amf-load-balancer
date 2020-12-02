package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceModifyResponseTransfer struct {
	DLNGUUPTNLInformation              *UPTransportLayerInformation                                              `aper:"valueLB:0,valueUB:1,optional"`
	ULNGUUPTNLInformation              *UPTransportLayerInformation                                              `aper:"valueLB:0,valueUB:1,optional"`
	QosFlowAddOrModifyResponseList     *QosFlowAddOrModifyResponseList                                           `aper:"optional"`
	AdditionalQosFlowPerTNLInformation *QosFlowPerTNLInformation                                                 `aper:"valueExt,optional"`
	QosFlowFailedToAddOrModifyList     *QosFlowList                                                              `aper:"optional"`
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceModifyResponseTransferExtIEs `aper:"optional"`
}
