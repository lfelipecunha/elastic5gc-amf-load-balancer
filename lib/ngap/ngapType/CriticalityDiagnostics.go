package ngapType

// Need to import "amfLoadBalancer/lib/aper" if it uses "aper"

type CriticalityDiagnostics struct {
	ProcedureCode             *ProcedureCode                                          `aper:"optional"`
	TriggeringMessage         *TriggeringMessage                                      `aper:"optional"`
	ProcedureCriticality      *Criticality                                            `aper:"optional"`
	IEsCriticalityDiagnostics *CriticalityDiagnosticsIEList                           `aper:"optional"`
	IEExtensions              *ProtocolExtensionContainerCriticalityDiagnosticsExtIEs `aper:"optional"`
}
