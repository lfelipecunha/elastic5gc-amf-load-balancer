package ngap

import (
	"amfLoadBalancer/lib/aper"
	"amfLoadBalancer/lib/ngap"
	"amfLoadBalancer/lib/ngap/ngapConvert"
	"amfLoadBalancer/lib/ngap/ngapType"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
	ngap_message "amfLoadBalancer/src/ngap/message"
	"encoding/hex"
	"net"
)

func HandleNGSetupRequest(ran *context.AmfRan, message *ngapType.NGAPPDU, originalMsg []byte) {
	var globalRANNodeID *ngapType.GlobalRANNodeID
	var rANNodeName *ngapType.RANNodeName
	var supportedTAList *ngapType.SupportedTAList
	var pagingDRX *ngapType.PagingDRX

	var cause ngapType.Cause

	if ran == nil {
		logger.NgapLog.Error("ran is nil")
		return
	}

	ran.NGSetupMsg = originalMsg

	if message == nil {
		logger.NgapLog.Error("NGAP Message is nil")
		return
	}

	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		logger.NgapLog.Error("Initiating Message is nil")
		return
	}
	nGSetupRequest := initiatingMessage.Value.NGSetupRequest
	if nGSetupRequest == nil {
		logger.NgapLog.Error("NGSetupRequest is nil")
		return
	}
	logger.NgapLog.Info("[AMF] NG Setup request")
	for i := 0; i < len(nGSetupRequest.ProtocolIEs.List); i++ {
		ie := nGSetupRequest.ProtocolIEs.List[i]
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDGlobalRANNodeID:
			globalRANNodeID = ie.Value.GlobalRANNodeID
			logger.NgapLog.Trace("[NGAP] Decode IE GlobalRANNodeID")
			if globalRANNodeID == nil {
				logger.NgapLog.Error("GlobalRANNodeID is nil")
				return
			}
		case ngapType.ProtocolIEIDSupportedTAList:
			supportedTAList = ie.Value.SupportedTAList
			logger.NgapLog.Trace("[NGAP] Decode IE SupportedTAList")
			if supportedTAList == nil {
				logger.NgapLog.Error("SupportedTAList is nil")
				return
			}
		case ngapType.ProtocolIEIDRANNodeName:
			rANNodeName = ie.Value.RANNodeName
			logger.NgapLog.Trace("[NGAP] Decode IE RANNodeName")
			if rANNodeName == nil {
				logger.NgapLog.Error("RANNodeName is nil")
				return
			}
		case ngapType.ProtocolIEIDDefaultPagingDRX:
			pagingDRX = ie.Value.DefaultPagingDRX
			logger.NgapLog.Trace("[NGAP] Decode IE DefaultPagingDRX")
			if pagingDRX == nil {
				logger.NgapLog.Error("DefaultPagingDRX is nil")
				return
			}
		}
	}

	ran.SetRanId(globalRANNodeID)
	if rANNodeName != nil {
		ran.Name = rANNodeName.Value
	}
	if pagingDRX != nil {
		logger.NgapLog.Tracef("PagingDRX[%d]", pagingDRX.Value)
	}

	for i := 0; i < len(supportedTAList.List); i++ {
		supportedTAItem := supportedTAList.List[i]
		tac := hex.EncodeToString(supportedTAItem.TAC.Value)
		capOfSupportTai := cap(ran.SupportedTAList)
		for j := 0; j < len(supportedTAItem.BroadcastPLMNList.List); j++ {
			supportedTAI := context.NewSupportedTAI()
			supportedTAI.Tai.Tac = tac
			broadcastPLMNItem := supportedTAItem.BroadcastPLMNList.List[j]
			plmnId := ngapConvert.PlmnIdToModels(broadcastPLMNItem.PLMNIdentity)
			supportedTAI.Tai.PlmnId = &plmnId
			capOfSNssaiList := cap(supportedTAI.SNssaiList)
			for k := 0; k < len(broadcastPLMNItem.TAISliceSupportList.List); k++ {
				tAISliceSupportItem := broadcastPLMNItem.TAISliceSupportList.List[k]
				if len(supportedTAI.SNssaiList) < capOfSNssaiList {
					supportedTAI.SNssaiList = append(supportedTAI.SNssaiList, ngapConvert.SNssaiToModels(tAISliceSupportItem.SNSSAI))
				} else {
					break
				}
			}
			logger.NgapLog.Tracef("PLMN_ID[MCC:%s MNC:%s] TAC[%s]", plmnId.Mcc, plmnId.Mnc, tac)
			if len(ran.SupportedTAList) < capOfSupportTai {
				ran.SupportedTAList = append(ran.SupportedTAList, supportedTAI)

			} else {
				break
			}
		}

	}

	if len(ran.SupportedTAList) == 0 {
		logger.NgapLog.Warn("NG-Setup failure: No supported TA exist in NG-Setup request")
		cause.Present = ngapType.CausePresentMisc
		cause.Misc = &ngapType.CauseMisc{
			Value: ngapType.CauseMiscPresentUnspecified,
		}
	} else {
		var found bool
		for i, tai := range ran.SupportedTAList {
			if context.InTaiList(tai.Tai, context.AMF_Self().SupportTaiLists) {
				logger.NgapLog.Tracef("SERVED_TAI_INDEX[%d]", i)
				found = true
				break
			}
		}
		if !found {
			logger.NgapLog.Warn("NG-Setup failure: Cannot find Served TAI in AMF")
			cause.Present = ngapType.CausePresentMisc
			cause.Misc = &ngapType.CauseMisc{
				Value: ngapType.CauseMiscPresentUnknownPLMN,
			}
		}
	}

	ngap_message.SendNGSetupResponse(ran)
}

func HandleInitialUEMessage(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var rANUENGAPID *ngapType.RANUENGAPID
	var ranUe *context.RanUe
	var err error

	if ran == nil {
		logger.NgapLog.Error("ran is nil")
		return
	}
	if message == nil {
		logger.NgapLog.Error("NGAP Message is nil")
		return
	}
	initiatingMessage := message.InitiatingMessage
	if initiatingMessage == nil {
		logger.NgapLog.Error("initiatingMessage is nil")
		return
	}
	initialUeMessage := initiatingMessage.Value.InitialUEMessage
	if initialUeMessage == nil {
		logger.NgapLog.Error("initialUeMessage is nil")
		return
	}

	logger.NgapLog.Info("[AMF] PDU Session Resource Setup Response")

	for _, ie := range initialUeMessage.ProtocolIEs.List {
		switch ie.Id.Value {
		case ngapType.ProtocolIEIDRANUENGAPID: // ignore
			rANUENGAPID = ie.Value.RANUENGAPID
			logger.NgapLog.Trace("[NGAP] Decode IE RanUeNgapID")
		}
	}

	printRanInfo(ran)

	if rANUENGAPID != nil {
		ranUe = ran.RanUeFindByRanUeNgapID(rANUENGAPID.Value)
		if ranUe == nil {
			ranUe, err = ran.NewRanUe(rANUENGAPID.Value)
			if err != nil {
				logger.NgapLog.Errorf("New RanUe Error: %+v", err)
				return
			}
		}
	}

	if ranUe != nil {
		ngSetupMessage, err := ngap_message.BuildNgSetupRequest(ranUe)
		ngap_message.SendToAmf(ran.Amf, ngSetupMessage)
		var returnMessage = make([]byte, 2048)
		var n int
		n, err = ran.Amf.Conn.Read(returnMessage)
		if err != nil {
			logger.NgapLog.Errorf("Error: %+v", err)
			return
		}
		logger.NgapLog.Debugf("Received: size[%i] message: %+v", n, returnMessage)

		curMessage, err := ngap.Encoder(*message)

		ngap_message.SendToAmf(ran.Amf, curMessage)

		_, err = ran.Amf.Conn.Read(returnMessage)
		if err != nil {
			logger.NgapLog.Errorf("Error: %+v", err)
			return
		}
		ngap_message.SendToRan(ranUe.Ran, returnMessage)
	}
	return
}

func printAndGetCause(cause *ngapType.Cause) (present int, value aper.Enumerated) {

	present = cause.Present
	switch cause.Present {
	case ngapType.CausePresentRadioNetwork:
		logger.NgapLog.Warnf("Cause RadioNetwork[%d]", cause.RadioNetwork.Value)
		value = cause.RadioNetwork.Value
	case ngapType.CausePresentTransport:
		logger.NgapLog.Warnf("Cause Transport[%d]", cause.Transport.Value)
		value = cause.Transport.Value
	case ngapType.CausePresentProtocol:
		logger.NgapLog.Warnf("Cause Protocol[%d]", cause.Protocol.Value)
		value = cause.Protocol.Value
	case ngapType.CausePresentNas:
		logger.NgapLog.Warnf("Cause Nas[%d]", cause.Nas.Value)
		value = cause.Nas.Value
	case ngapType.CausePresentMisc:
		logger.NgapLog.Warnf("Cause Misc[%d]", cause.Misc.Value)
		value = cause.Misc.Value
	default:
		logger.NgapLog.Errorf("Invalid Cause group[%d]", cause.Present)
	}
	return
}

func printCriticalityDiagnostics(criticalityDiagnostics *ngapType.CriticalityDiagnostics) {
	logger.NgapLog.Trace("Criticality Diagnostics")

	if criticalityDiagnostics.ProcedureCriticality != nil {
		switch criticalityDiagnostics.ProcedureCriticality.Value {
		case ngapType.CriticalityPresentReject:
			logger.NgapLog.Trace("Procedure Criticality: Reject")
		case ngapType.CriticalityPresentIgnore:
			logger.NgapLog.Trace("Procedure Criticality: Ignore")
		case ngapType.CriticalityPresentNotify:
			logger.NgapLog.Trace("Procedure Criticality: Notify")
		}
	}

	if criticalityDiagnostics.IEsCriticalityDiagnostics != nil {
		for _, ieCriticalityDiagnostics := range criticalityDiagnostics.IEsCriticalityDiagnostics.List {
			logger.NgapLog.Tracef("IE ID: %d", ieCriticalityDiagnostics.IEID.Value)

			switch ieCriticalityDiagnostics.IECriticality.Value {
			case ngapType.CriticalityPresentReject:
				logger.NgapLog.Trace("Criticality Reject")
			case ngapType.CriticalityPresentNotify:
				logger.NgapLog.Trace("Criticality Notify")
			}

			switch ieCriticalityDiagnostics.TypeOfError.Value {
			case ngapType.TypeOfErrorPresentNotUnderstood:
				logger.NgapLog.Trace("Type of error: Not understood")
			case ngapType.TypeOfErrorPresentMissing:
				logger.NgapLog.Trace("Type of error: Missing")
			}
		}
	}
}

func buildCriticalityDiagnostics(
	procedureCode *int64,
	triggeringMessage *aper.Enumerated,
	procedureCriticality *aper.Enumerated,
	iesCriticalityDiagnostics *ngapType.CriticalityDiagnosticsIEList) (
	criticalityDiagnostics ngapType.CriticalityDiagnostics) {

	if procedureCode != nil {
		criticalityDiagnostics.ProcedureCode = new(ngapType.ProcedureCode)
		criticalityDiagnostics.ProcedureCode.Value = *procedureCode
	}

	if triggeringMessage != nil {
		criticalityDiagnostics.TriggeringMessage = new(ngapType.TriggeringMessage)
		criticalityDiagnostics.TriggeringMessage.Value = *triggeringMessage
	}

	if procedureCriticality != nil {
		criticalityDiagnostics.ProcedureCriticality = new(ngapType.Criticality)
		criticalityDiagnostics.ProcedureCriticality.Value = *procedureCriticality
	}

	if iesCriticalityDiagnostics != nil {
		criticalityDiagnostics.IEsCriticalityDiagnostics = iesCriticalityDiagnostics
	}

	return criticalityDiagnostics
}

func buildCriticalityDiagnosticsIEItem(ieCriticality aper.Enumerated, ieID int64, typeOfErr aper.Enumerated) (
	item ngapType.CriticalityDiagnosticsIEItem) {

	item = ngapType.CriticalityDiagnosticsIEItem{
		IECriticality: ngapType.Criticality{
			Value: ieCriticality,
		},
		IEID: ngapType.ProtocolIEID{
			Value: ieID,
		},
		TypeOfError: ngapType.TypeOfError{
			Value: typeOfErr,
		},
	}

	return item
}

func printRanInfo(ran *context.AmfRan) {
	var addr net.Addr
	var remote_ip string

	addr = ran.Conn.RemoteAddr()

	remote_ip = "Undefined"
	if addr != nil {
		remote_ip = addr.String()
	}
	switch ran.RanPresent {
	case context.RanPresentGNbId:
		logger.NgapLog.Tracef("IP[%s] GNbId[%s]", remote_ip, ran.RanId.GNbId.GNBValue)
	case context.RanPresentNgeNbId:
		logger.NgapLog.Tracef("IP[%s] NgeNbId[%s]", remote_ip, ran.RanId.NgeNbId)
	case context.RanPresentN3IwfId:
		logger.NgapLog.Tracef("IP[%s] N3IwfId[%s]", remote_ip, ran.RanId.N3IwfId)
	}
}
