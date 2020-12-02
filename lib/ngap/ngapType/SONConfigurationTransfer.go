package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type SONConfigurationTransfer struct {
	TargetRANNodeID        TargetRANNodeID                                           `aper:"valueExt"`
	SourceRANNodeID        SourceRANNodeID                                           `aper:"valueExt"`
	SONInformation         SONInformation                                            `aper:"valueLB:0,valueUB:2"`
	XnTNLConfigurationInfo XnTNLConfigurationInfo                                    `aper:"valueExt"`
	IEExtensions           *ProtocolExtensionContainerSONConfigurationTransferExtIEs `aper:"optional"`
}
