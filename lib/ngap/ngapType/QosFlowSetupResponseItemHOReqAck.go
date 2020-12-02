package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type QosFlowSetupResponseItemHOReqAck struct {
	QosFlowIdentifier      QosFlowIdentifier
	DataForwardingAccepted *DataForwardingAccepted                                           `aper:"optional"`
	IEExtensions           *ProtocolExtensionContainerQosFlowSetupResponseItemHOReqAckExtIEs `aper:"optional"`
}
