package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	AMFPagingTargetPresentNothing int = iota /* No components present */
	AMFPagingTargetPresentGlobalRANNodeID
	AMFPagingTargetPresentTAI
	AMFPagingTargetPresentChoiceExtensions
)

type AMFPagingTarget struct {
	Present          int
	GlobalRANNodeID  *GlobalRANNodeID `aper:"valueLB:0,valueUB:3"`
	TAI              *TAI             `aper:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerAMFPagingTargetExtIEs
}
