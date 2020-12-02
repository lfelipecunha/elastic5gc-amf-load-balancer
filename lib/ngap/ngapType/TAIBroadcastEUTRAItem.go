package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TAIBroadcastEUTRAItem struct {
	TAI                      TAI `aper:"valueExt"`
	CompletedCellsInTAIEUTRA CompletedCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAIBroadcastEUTRAItemExtIEs `aper:"optional"`
}
