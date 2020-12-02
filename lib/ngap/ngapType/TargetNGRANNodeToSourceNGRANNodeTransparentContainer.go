package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type TargetNGRANNodeToSourceNGRANNodeTransparentContainer struct {
	RRCContainer RRCContainer
	IEExtensions *ProtocolExtensionContainerTargetNGRANNodeToSourceNGRANNodeTransparentContainerExtIEs `aper:"optional"`
}
