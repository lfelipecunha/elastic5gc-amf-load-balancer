package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UserLocationInformationN3IWF struct {
	IPAddress    TransportLayerAddress
	PortNumber   PortNumber
	IEExtensions *ProtocolExtensionContainerUserLocationInformationN3IWFExtIEs `aper:"optional"`
}
