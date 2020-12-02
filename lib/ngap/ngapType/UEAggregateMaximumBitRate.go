package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UEAggregateMaximumBitRate struct {
	UEAggregateMaximumBitRateDL BitRate
	UEAggregateMaximumBitRateUL BitRate
	IEExtensions                *ProtocolExtensionContainerUEAggregateMaximumBitRateExtIEs `aper:"optional"`
}
