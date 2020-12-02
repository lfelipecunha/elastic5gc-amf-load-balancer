package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type DRBStatusDL12 struct {
	DLCOUNTValue COUNTValueForPDCPSN12                          `aper:"valueExt"`
	IEExtension  *ProtocolExtensionContainerDRBStatusDL12ExtIEs `aper:"optional"`
}
