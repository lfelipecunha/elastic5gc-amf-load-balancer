package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PDUSessionResourceSetupUnsuccessfulTransfer struct {
	Cause                  Cause                                                                        `aper:"valueLB:0,valueUB:5"`
	CriticalityDiagnostics *CriticalityDiagnostics                                                      `aper:"valueExt,optional"`
	IEExtensions           *ProtocolExtensionContainerPDUSessionResourceSetupUnsuccessfulTransferExtIEs `aper:"optional"`
}
