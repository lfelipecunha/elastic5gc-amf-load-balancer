package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type AMFTNLAssociationToUpdateItem struct {
	AMFTNLAssociationAddress CPTransportLayerInformation                                    `aper:"valueLB:0,valueUB:1"`
	TNLAssociationUsage      *TNLAssociationUsage                                           `aper:"optional"`
	TNLAddressWeightFactor   *TNLAddressWeightFactor                                        `aper:"optional"`
	IEExtensions             *ProtocolExtensionContainerAMFTNLAssociationToUpdateItemExtIEs `aper:"optional"`
}
