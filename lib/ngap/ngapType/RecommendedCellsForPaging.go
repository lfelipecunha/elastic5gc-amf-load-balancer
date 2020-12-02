package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type RecommendedCellsForPaging struct {
	RecommendedCellList RecommendedCellList
	IEExtensions        *ProtocolExtensionContainerRecommendedCellsForPagingExtIEs `aper:"optional"`
}
