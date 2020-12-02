package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type QosFlowSetupResponseItemSURes struct {
	QosFlowIdentifier QosFlowIdentifier
	IEExtensions      *ProtocolExtensionContainerQosFlowSetupResponseItemSUResExtIEs `aper:"optional"`
}
