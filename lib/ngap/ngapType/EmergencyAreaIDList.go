package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct EmergencyAreaIDList */
/* EmergencyAreaID */
type EmergencyAreaIDList struct {
	List []EmergencyAreaID `aper:"sizeLB:1,sizeUB:65535"`
}
