package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

/* Sequence of = 35, FULL Name = struct PDUSessionResourceModifyListModCfm */
/* PDUSessionResourceModifyItemModCfm */
type PDUSessionResourceModifyListModCfm struct {
	List []PDUSessionResourceModifyItemModCfm `aper:"valueExt,sizeLB:1,sizeUB:256"`
}
