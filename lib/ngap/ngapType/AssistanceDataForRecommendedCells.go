package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type AssistanceDataForRecommendedCells struct {
	RecommendedCellsForPaging RecommendedCellsForPaging                                          `aper:"valueExt"`
	IEExtensions              *ProtocolExtensionContainerAssistanceDataForRecommendedCellsExtIEs `aper:"optional"`
}
