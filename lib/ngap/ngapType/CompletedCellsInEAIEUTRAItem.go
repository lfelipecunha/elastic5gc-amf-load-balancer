package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type CompletedCellsInEAIEUTRAItem struct {
	EUTRACGI     EUTRACGI                                                      `aper:"valueExt"`
	IEExtensions *ProtocolExtensionContainerCompletedCellsInEAIEUTRAItemExtIEs `aper:"optional"`
}
