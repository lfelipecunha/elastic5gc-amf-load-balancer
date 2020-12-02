package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UEPresenceInAreaOfInterestItem struct {
	LocationReportingReferenceID LocationReportingReferenceID
	UEPresence                   UEPresence
	IEExtensions                 *ProtocolExtensionContainerUEPresenceInAreaOfInterestItemExtIEs `aper:"optional"`
}
