package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type LastVisitedCellItem struct {
	LastVisitedCellInformation LastVisitedCellInformation                           `aper:"valueLB:0,valueUB:4"`
	IEExtensions               *ProtocolExtensionContainerLastVisitedCellItemExtIEs `aper:"optional"`
}
