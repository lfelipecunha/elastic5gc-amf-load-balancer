package ngap

import (
	"amfLoadBalancer/lib/aper"
	"amfLoadBalancer/lib/nas"
	"amfLoadBalancer/lib/ngap"
	"amfLoadBalancer/lib/ngap/ngapConvert"
	"amfLoadBalancer/lib/ngap/ngapType"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
	ngap_message "amfLoadBalancer/src/ngap/message"
	"encoding/hex"
	"net"
	"reflect"
)

// HandleNGSetupRequest Handle NG Setup Request
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
			plmnID := ngapConvert.PlmnIdToModels(broadcastPLMNItem.PLMNIdentity)
			supportedTAI.Tai.PlmnId = &plmnID
			capOfSNssaiList := cap(supportedTAI.SNssaiList)
			for k := 0; k < len(broadcastPLMNItem.TAISliceSupportList.List); k++ {
				tAISliceSupportItem := broadcastPLMNItem.TAISliceSupportList.List[k]
				if len(supportedTAI.SNssaiList) < capOfSNssaiList {
					supportedTAI.SNssaiList = append(supportedTAI.SNssaiList, ngapConvert.SNssaiToModels(tAISliceSupportItem.SNSSAI))
				} else {
					break
				}
			}
			logger.NgapLog.Tracef("PLMN_ID[MCC:%s MNC:%s] TAC[%s]", plmnID.Mcc, plmnID.Mnc, tac)
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

// HandleInitialUEMessage Handle Initial UE Message
func HandleInitialUEMessage(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var rANUENGAPID *ngapType.RANUENGAPID
	var ranUe *context.RanUe
	if ran == nil {
		logger.NgapLog.Error("ran is nil")
		return
	}
	if message == nil {
		logger.NgapLog.Error("NGAP Message is nil")
		return
	}

	rANUENGAPID = GetRanUeID(message)
	logger.NgapLog.Debugf("[NGAP] Decode IE RanUeNgapID [%i]", rANUENGAPID.Value)

	printRanInfo(ran)

	if rANUENGAPID != nil {
		ranUe = GetRanUe(ran, rANUENGAPID, true)
	}

	if ranUe != nil {
		ngSetupMessage, err := ngap_message.BuildNgSetupRequest(ranUe)

		balancer := context.AMF_Self().Balancer
		amf, err := ranUe.Ran.GetConnToAmf(balancer.SelectAmf(ranUe))

		if err != nil {
			logger.NgapLog.Errorf("Error on connect to AMF: %+v", err)
			return
		}

		ngap_message.SendToAmf(amf, ngSetupMessage)
		var returnMessage = make([]byte, 2048)

		_, err = amf.Conn.Read(returnMessage)
		if err != nil {
			logger.NgapLog.Errorf("Error: %+v", err)
			return
		}

		ProxyMessage(ranUe, amf, message, true)
	}
	return
}

func HandleGenericMessages(ran *context.AmfRan, message *ngapType.NGAPPDU) {
	var rANUENGAPID *ngapType.RANUENGAPID
	var ranUe *context.RanUe
	var procedureCode int64
	var read bool

	if ran == nil {
		logger.NgapLog.Error("ran is nil")
		return
	}
	if message == nil {
		logger.NgapLog.Error("NGAP Message is nil")
		return
	}

	if message.InitiatingMessage != nil {
		procedureCode = message.InitiatingMessage.ProcedureCode.Value
	} else if message.SuccessfulOutcome != nil {
		procedureCode = message.SuccessfulOutcome.ProcedureCode.Value
	} else if message.UnsuccessfulOutcome != nil {
		procedureCode = message.UnsuccessfulOutcome.ProcedureCode.Value
	}
	logger.NgapLog.Infof("Handling %d message", procedureCode)

	rANUENGAPID = GetRanUeID(message)
	logger.NgapLog.Debugf("Decode IE RanUeNgapID [%i]", rANUENGAPID.Value)

	printRanInfo(ran)

	switch procedureCode {
	case ngapType.ProcedureCodeInitialContextSetup,
		ngapType.ProcedureCodePDUSessionResourceSetup:
		read = false
	case ngapType.ProcedureCodeUplinkNASTransport:
		nasProcCode := getNasProcedureCode(message.InitiatingMessage.Value.UplinkNASTransport.ProtocolIEs.List)
		logger.NgapLog.Debugf("NAS CODE [%d]", nasProcCode)
		switch nasProcCode {
		case nas.MsgTypeRegistrationComplete:
			read = false
		default:
			read = true
		}
	default:
		read = true
	}

	if rANUENGAPID != nil {
		ranUe = GetRanUe(ran, rANUENGAPID, false)
		if ranUe != nil {
			amf, err := ranUe.Ran.GetConnToAmf(context.AMF_Self().Balancer.SelectAmf(ranUe))
			if err != nil {
				logger.NgapLog.Errorf("Error on connect to AMF: %+v", err)
				return
			}
			ProxyMessage(ranUe, amf, message, read)
		}
	}
	return
}

func getNasProcedureCode(list []ngapType.UplinkNASTransportIEs) uint8 {
	var payload []byte
	for _, item := range list {
		if item.Id.Value == ngapType.ProtocolIEIDNASPDU {
			payload = item.Value.NASPDU.Value
			msg := new(nas.Message)
			msg.SecurityHeaderType = uint8(nas.GetSecurityHeaderType(payload) & 0x0f)
			if msg.SecurityHeaderType != nas.SecurityHeaderTypePlainNas {
				payload = payload[7:]
			}
			msg.PlainNasDecode(&payload)
			return msg.GmmMessage.GetMessageType()
		}
	}
	return 0
}

func ProxyMessage(ranUe *context.RanUe, amf *context.Amf, message *ngapType.NGAPPDU, read bool) {
	var returnMessage = make([]byte, 2048)
	curMessage, err := ngap.Encoder(*message)
	var procedureCode int64
	if message.InitiatingMessage != nil {
		procedureCode = message.InitiatingMessage.ProcedureCode.Value
	} else if message.SuccessfulOutcome != nil {
		procedureCode = message.SuccessfulOutcome.ProcedureCode.Value
	} else if message.UnsuccessfulOutcome != nil {
		procedureCode = message.UnsuccessfulOutcome.ProcedureCode.Value
	}

	logger.NgapLog.Infof("Send original message (procedure: %d) to AMF[%s]", procedureCode, amf.AmfData.IP)
	ngap_message.SendToAmf(amf, curMessage)
	if read {
		logger.NgapLog.Debugf("Procedure[%d] Waiting AMF response...", procedureCode)
		_, err = amf.Conn.Read(returnMessage)
		logger.NgapLog.Debugf("Procedure[%d] Received AMF response", procedureCode)
		if err != nil {
			logger.NgapLog.Errorf("Error: %+v", err)
			return
		}
		logger.NgapLog.Infof("Send returned message from AMF[%s] to RAN", amf.AmfData.IP)
		ngap_message.SendToRan(ranUe.Ran, returnMessage)
	} else {
		logger.NgapLog.Info("This message does not return data")
	}
}

// GetRanUe Get context.RanUE object based on UE NGAP ID
func GetRanUe(ran *context.AmfRan, ranUEID *ngapType.RANUENGAPID, create bool) *context.RanUe {
	var err error
	ranUe := ran.RanUeFindByRanUeNgapID(ranUEID.Value)
	if ranUe == nil && create {
		ranUe, err = ran.NewRanUe(ranUEID.Value)
		if err != nil {
			logger.NgapLog.Errorf("New RanUe Error: %+v", err)
		}
	}
	return ranUe
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

// printRanInfo Print Ran Info
func printRanInfo(ran *context.AmfRan) {
	var addr net.Addr
	var remoteIP string

	addr = ran.Conn.RemoteAddr()

	remoteIP = "Undefined"
	if addr != nil {
		remoteIP = addr.String()
	}
	switch ran.RanPresent {
	case context.RanPresentGNbId:
		logger.NgapLog.Tracef("IP[%s] GNbId[%s]", remoteIP, ran.RanId.GNbId.GNBValue)
	case context.RanPresentNgeNbId:
		logger.NgapLog.Tracef("IP[%s] NgeNbId[%s]", remoteIP, ran.RanId.NgeNbId)
	case context.RanPresentN3IwfId:
		logger.NgapLog.Tracef("IP[%s] N3IwfId[%s]", remoteIP, ran.RanId.N3IwfId)
	}
}

// GetRanUeID get NGAP ID for UE in RAN context
func GetRanUeID(message *ngapType.NGAPPDU) *ngapType.RANUENGAPID {
	if message == nil {
		logger.NgapLog.Error("NGAP Message is nil")
		return nil
	}

	var r reflect.Value
	switch {
	case message.InitiatingMessage != nil:
		r = reflect.ValueOf(message.InitiatingMessage.Value)
	case message.SuccessfulOutcome != nil:
		r = reflect.ValueOf(message.SuccessfulOutcome.Value)
	case message.UnsuccessfulOutcome != nil:
		r = reflect.ValueOf(message.UnsuccessfulOutcome.Value)
	}

	for i := 0; i < r.NumField(); i++ {
		field := r.Field(i)
		if field.Kind() == reflect.Ptr {
			field = field.Elem()
			if field.IsValid() {
				return searchID(field.Field(0).Field(0))
			}
		}
	}

	return nil
}

func searchID(list reflect.Value) *ngapType.RANUENGAPID {
	var i int
	for i = 0; i < list.Len(); i++ {
		item := list.Index(i)
		if item.FieldByName("Id").FieldByName("Value").Int() == ngapType.ProtocolIEIDRANUENGAPID {
			return item.FieldByName("Value").FieldByName("RANUENGAPID").Interface().(*ngapType.RANUENGAPID)
		}
	}
	return nil
}
