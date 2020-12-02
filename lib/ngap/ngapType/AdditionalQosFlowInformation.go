package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	AdditionalQosFlowInformationPresentMoreLikely aper.Enumerated = 0
)

type AdditionalQosFlowInformation struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:0"`
}
