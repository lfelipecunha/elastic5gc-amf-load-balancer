package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type GlobalGNBID struct {
	PLMNIdentity PLMNIdentity
	GNBID        GNBID                                        `aper:"valueLB:0,valueUB:1"`
	IEExtensions *ProtocolExtensionContainerGlobalGNBIDExtIEs `aper:"optional"`
}
