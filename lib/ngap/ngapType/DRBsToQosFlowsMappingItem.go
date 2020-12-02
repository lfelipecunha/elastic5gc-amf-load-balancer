package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type DRBsToQosFlowsMappingItem struct {
	DRBID                 DRBID
	AssociatedQosFlowList AssociatedQosFlowList
	IEExtensions          *ProtocolExtensionContainerDRBsToQosFlowsMappingItemExtIEs `aper:"optional"`
}
