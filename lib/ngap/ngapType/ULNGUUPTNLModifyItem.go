package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type ULNGUUPTNLModifyItem struct {
	ULNGUUPTNLInformation UPTransportLayerInformation                           `aper:"valueLB:0,valueUB:1"`
	DLNGUUPTNLInformation UPTransportLayerInformation                           `aper:"valueLB:0,valueUB:1"`
	IEExtensions          *ProtocolExtensionContainerULNGUUPTNLModifyItemExtIEs `aper:"optional"`
}
