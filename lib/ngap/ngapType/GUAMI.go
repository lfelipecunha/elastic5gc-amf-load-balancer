package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type GUAMI struct {
	PLMNIdentity PLMNIdentity
	AMFRegionID  AMFRegionID
	AMFSetID     AMFSetID
	AMFPointer   AMFPointer
	IEExtensions *ProtocolExtensionContainerGUAMIExtIEs `aper:"optional"`
}
