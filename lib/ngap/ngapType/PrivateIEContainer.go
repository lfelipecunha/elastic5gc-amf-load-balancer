package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PrivateIE_Container_6516P0 */
/* PrivateMessageIEs */
type PrivateIEContainerPrivateMessageIEs struct {
	List []PrivateMessageIEs `aper:"sizeLB:1,sizeUB:65535"`
}
