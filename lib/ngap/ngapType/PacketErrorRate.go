package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type PacketErrorRate struct {
	PERScalar    int64                                            `aper:"valueExt,valueLB:0,valueUB:9"`
	PERExponent  int64                                            `aper:"valueExt,valueLB:0,valueUB:9"`
	IEExtensions *ProtocolExtensionContainerPacketErrorRateExtIEs `aper:"optional"`
}
