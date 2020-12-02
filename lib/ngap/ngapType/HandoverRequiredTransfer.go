package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type HandoverRequiredTransfer struct {
	DirectForwardingPathAvailability *DirectForwardingPathAvailability                         `aper:"optional"`
	IEExtensions                     *ProtocolExtensionContainerHandoverRequiredTransferExtIEs `aper:"optional"`
}
