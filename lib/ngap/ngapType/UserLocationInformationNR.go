package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UserLocationInformationNR struct {
	NRCGI        NRCGI                                                      `aper:"valueExt"`
	TAI          TAI                                                        `aper:"valueExt"`
	TimeStamp    *TimeStamp                                                 `aper:"optional"`
	IEExtensions *ProtocolExtensionContainerUserLocationInformationNRExtIEs `aper:"optional"`
}
