package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type LocationReportingRequestType struct {
	EventType                                 EventType
	ReportArea                                ReportArea
	AreaOfInterestList                        *AreaOfInterestList                                           `aper:"optional"`
	LocationReportingReferenceIDToBeCancelled *LocationReportingReferenceID                                 `aper:"optional"`
	IEExtensions                              *ProtocolExtensionContainerLocationReportingRequestTypeExtIEs `aper:"optional"`
}
