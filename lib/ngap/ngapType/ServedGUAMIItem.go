package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type ServedGUAMIItem struct {
	GUAMI         GUAMI                                            `aper:"valueExt"`
	BackupAMFName *AMFName                                         `aper:"sizeExt,sizeLB:1,sizeUB:150,optional"`
	IEExtensions  *ProtocolExtensionContainerServedGUAMIItemExtIEs `aper:"optional"`
}
