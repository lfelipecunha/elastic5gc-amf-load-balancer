package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceModifyItemModRes struct {
	PDUSessionID                             PDUSessionID
	PDUSessionResourceModifyResponseTransfer *aper.OctetString                                                   `aper:"optional"`
	IEExtensions                             *ProtocolExtensionContainerPDUSessionResourceModifyItemModResExtIEs `aper:"optional"`
}
