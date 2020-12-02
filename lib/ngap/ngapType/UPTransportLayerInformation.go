package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	UPTransportLayerInformationPresentNothing int = iota /* No components present */
	UPTransportLayerInformationPresentGTPTunnel
	UPTransportLayerInformationPresentChoiceExtensions
)

type UPTransportLayerInformation struct {
	Present          int
	GTPTunnel        *GTPTunnel `aper:"valueExt"`
	ChoiceExtensions *ProtocolIESingleContainerUPTransportLayerInformationExtIEs
}
