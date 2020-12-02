package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type CellType struct {
	CellSize     CellSize
	IEExtensions *ProtocolExtensionContainerCellTypeExtIEs `aper:"optional"`
}
