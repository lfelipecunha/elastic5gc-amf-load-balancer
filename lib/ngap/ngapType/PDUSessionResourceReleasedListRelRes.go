package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceReleasedListRelRes */
/* PDUSessionResourceReleasedItemRelRes */
type PDUSessionResourceReleasedListRelRes struct {
	List []PDUSessionResourceReleasedItemRelRes `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
