package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PathSwitchRequestAcknowledgeTransfer struct {
	ULNGUUPTNLInformation *UPTransportLayerInformation                                          `aper:"valueLB:0,valueUB:1,optional"`
	SecurityIndication    *SecurityIndication                                                   `aper:"valueExt,optional"`
	IEExtensions          *ProtocolExtensionContainerPathSwitchRequestAcknowledgeTransferExtIEs `aper:"optional"`
}
