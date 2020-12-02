package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	UPTNLInformationPresentNothing int = iota /* No components present */
	UPTNLInformationPresentSingleTNLInformation
	UPTNLInformationPresentMultipleTNLInformation
	UPTNLInformationPresentChoiceExtensions
)

type UPTNLInformation struct {
	Present                int
	SingleTNLInformation   *SingleTNLInformation   `aper:"valueExt"`
	MultipleTNLInformation *MultipleTNLInformation `aper:"valueExt"`
	ChoiceExtensions       *ProtocolIESingleContainerUPTNLInformationExtIEs
}
