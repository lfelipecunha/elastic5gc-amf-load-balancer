package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceAdmittedItem struct {
	PDUSessionID                       PDUSessionID
	HandoverRequestAcknowledgeTransfer aper.OctetString
	IEExtensions                       *ProtocolExtensionContainerPDUSessionResourceAdmittedItemExtIEs `aper:"optional"`
}
