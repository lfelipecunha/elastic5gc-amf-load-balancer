package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceModifyIndicationTransfer struct {
	DLUPTNLInformation *UPTNLInformation                                                           `aper:"valueLB:0,valueUB:2,optional"`
	IEExtensions       *ProtocolExtensionContainerPDUSessionResourceModifyIndicationTransferExtIEs `aper:"optional"`
}
