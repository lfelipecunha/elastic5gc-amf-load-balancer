package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

const (
	EventTypePresentDirect                          aper.Enumerated = 0
	EventTypePresentChangeOfServeCell               aper.Enumerated = 1
	EventTypePresentUePresenceInAreaOfInterest      aper.Enumerated = 2
	EventTypePresentStopChangeOfServeCell           aper.Enumerated = 3
	EventTypePresentStopUePresenceInAreaOfInterest  aper.Enumerated = 4
	EventTypePresentCancelLocationReportingForTheUe aper.Enumerated = 5
)

type EventType struct {
	Value aper.Enumerated `aper:"valueExt,valueLB:0,valueUB:5"`
}
