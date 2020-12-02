package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	MaximumIntegrityProtectedDataRatePresentBitrate64kbs  aper.Enumerated = 0
	MaximumIntegrityProtectedDataRatePresentMaximumUERate aper.Enumerated = 1
)

type MaximumIntegrityProtectedDataRate struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:1"`
}
