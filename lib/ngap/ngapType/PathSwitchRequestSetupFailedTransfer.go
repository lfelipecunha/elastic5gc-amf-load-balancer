package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PathSwitchRequestSetupFailedTransfer struct {
	Cause        Cause                                                                 `aper:"valueLB:0,valueUB:5"`
	IEExtensions *ProtocolExtensionContainerPathSwitchRequestSetupFailedTransferExtIEs `aper:"optional"`
}
