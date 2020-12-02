package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TAICancelledEUTRAItem struct {
	TAI                      TAI `aper:"valueExt"`
	CancelledCellsInTAIEUTRA CancelledCellsInTAIEUTRA
	IEExtensions             *ProtocolExtensionContainerTAICancelledEUTRAItemExtIEs `aper:"optional"`
}
