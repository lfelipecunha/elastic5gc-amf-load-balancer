package ngapType

import "amfLoadBalancer/lib/aper"

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type NRencryptionAlgorithms struct {
	Value aper.BitString `aper:"sizeExt,sizeLB:16,sizeUB:16"`
}
