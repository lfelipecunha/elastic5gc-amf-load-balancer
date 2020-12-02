package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	GNBIDPresentNothing int = iota /* No components present */
	GNBIDPresentGNBID
	GNBIDPresentChoiceExtensions
)

type GNBID struct {
	Present          int
	GNBID            *aper.BitString `aper:"sizeLB:22,sizeUB:32"`
	ChoiceExtensions *ProtocolIESingleContainerGNBIDExtIEs
}
