package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UEAssociatedLogicalNGConnectionItem struct {
	AMFUENGAPID  *AMFUENGAPID                                                         `aper:"optional"`
	RANUENGAPID  *RANUENGAPID                                                         `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUEAssociatedLogicalNGConnectionItemExtIEs `aper:"optional"`
}
