package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	UEPagingIdentityPresentNothing int = iota /* No components present */
	UEPagingIdentityPresentFiveGSTMSI
	UEPagingIdentityPresentChoiceExtensions
)

type UEPagingIdentity struct {
	Present          int
	FiveGSTMSI       *FiveGSTMSI `aper:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerUEPagingIdentityExtIEs
}
