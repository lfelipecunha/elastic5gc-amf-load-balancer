package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct RATRestrictions */
/* RATRestrictionsItem */
type RATRestrictions struct {
	List []RATRestrictionsItem `aper:"valueExt,sizeLB:0,sizeUB:16"`
}
