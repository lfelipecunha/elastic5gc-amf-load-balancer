package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type UnavailableGUAMIItem struct {
	GUAMI                        GUAMI                                                 `aper:"valueExt"`
	TimerApproachForGUAMIRemoval *TimerApproachForGUAMIRemoval                         `aper:"optional"`
	BackupAMFName                *AMFName                                              `aper:"sizeExt,sizeLB:1,sizeUB:150,optional"`
	IEExtensions                 *ProtocolExtensionContainerUnavailableGUAMIItemExtIEs `aper:"optional"`
}
