package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type SupportedTAItem struct {
	TAC               TAC
	BroadcastPLMNList BroadcastPLMNList
	IEExtensions      *ProtocolExtensionContainerSupportedTAItemExtIEs `aper:"optional"`
}
