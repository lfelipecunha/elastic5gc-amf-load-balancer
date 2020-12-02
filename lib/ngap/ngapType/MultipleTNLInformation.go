package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type MultipleTNLInformation struct {
	TNLInformationList TNLInformationList
	IEExtensions       *ProtocolExtensionContainerMultipleTNLInformationExtIEs `aper:"optional"`
}
