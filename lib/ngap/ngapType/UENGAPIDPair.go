package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UENGAPIDPair struct {
	AMFUENGAPID  AMFUENGAPID
	RANUENGAPID  RANUENGAPID
	IEExtensions *ProtocolExtensionContainerUENGAPIDPairExtIEs `aper:"optional"`
}
