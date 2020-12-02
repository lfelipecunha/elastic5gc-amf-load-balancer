package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TAIListForPagingItem struct {
	TAI          TAI                                                   `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerTAIListForPagingItemExtIEs `aper:"optional"`
}
