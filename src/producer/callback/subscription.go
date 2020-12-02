package callback

import (
	"context"
	"amfLoadBalancer/lib/openapi/Namf_Communication"
	"amfLoadBalancer/lib/openapi/models"
	amf_context "amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
	"reflect"
)

func SendAmfStatusChangeNotify(amfStatus string, guamiList []models.Guami) {
	amfSelf := amf_context.AMF_Self()

	amfSelf.AMFStatusSubscriptions.Range(func(key, value interface{}) bool {
		subscriptionData := value.(models.SubscriptionData)

		configuration := Namf_Communication.NewConfiguration()
		client := Namf_Communication.NewAPIClient(configuration)
		amfStatusNotification := models.AmfStatusChangeNotification{}
		var amfStatusInfo = models.AmfStatusInfo{}

		for _, guami := range guamiList {
			for _, subGumi := range subscriptionData.GuamiList {
				if reflect.DeepEqual(guami, subGumi) {
					//AMF status is available
					amfStatusInfo.GuamiList = append(amfStatusInfo.GuamiList, guami)
				}
			}
		}

		amfStatusInfo = models.AmfStatusInfo{
			StatusChange:     (models.StatusChange)(amfStatus),
			TargetAmfRemoval: "",
			TargetAmfFailure: "",
		}

		amfStatusNotification.AmfStatusInfoList = append(amfStatusNotification.AmfStatusInfoList, amfStatusInfo)
		uri := subscriptionData.AmfStatusUri

		logger.ProducerLog.Infof("[AMF] Send Amf Status Change Notify to %s", uri)
		httpResponse, err := client.AmfStatusChangeCallbackDocumentApiServiceCallbackDocumentApi.
			AmfStatusChangeNotify(context.Background(), uri, amfStatusNotification)
		if err != nil {
			if httpResponse == nil {
				HttpLog.Errorln(err.Error())
			} else if err.Error() != httpResponse.Status {
				HttpLog.Errorln(err.Error())
			}
		}
		return true
	})
}
