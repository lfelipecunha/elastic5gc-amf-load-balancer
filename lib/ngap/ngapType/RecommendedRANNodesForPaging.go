package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type RecommendedRANNodesForPaging struct {
	RecommendedRANNodeList RecommendedRANNodeList
	IEExtensions           *ProtocolExtensionContainerRecommendedRANNodesForPagingExtIEs `aper:"optional"`
}
