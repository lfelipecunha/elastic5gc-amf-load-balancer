package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type AllowedNSSAIItem struct {
	SNSSAI       SNSSAI                                            `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerAllowedNSSAIItemExtIEs `aper:"optional"`
}
