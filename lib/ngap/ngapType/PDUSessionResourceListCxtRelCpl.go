package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceListCxtRelCpl */
/* PDUSessionResourceItemCxtRelCpl */
type PDUSessionResourceListCxtRelCpl struct {
	List []PDUSessionResourceItemCxtRelCpl `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
