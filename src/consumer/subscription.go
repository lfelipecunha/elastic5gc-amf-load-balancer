package consumer

import (
	"amfLoadBalancer/lib/openapi/Nnrf_NFManagement"
	"amfLoadBalancer/lib/openapi/models"
	"amfLoadBalancer/src/logger"
	"context"
	"net/http"
)

func SubscribeAmfsChanges(nrfUri string, callbackUri string) {
	configuration := Nnrf_NFManagement.NewConfiguration()
	configuration.SetBasePath(nrfUri)
	client := Nnrf_NFManagement.NewAPIClient(configuration)
	data := models.NrfSubscriptionData{NfStatusNotificationUri: callbackUri, SubscrCond: models.ServiceNameCond{ServiceName: "namf-evts"}}
	_, httpResponse, err := client.SubscriptionsCollectionApi.CreateSubscription(context.TODO(), data)
	if err != nil {
		logger.ConsumerLog.Errorf("Error on subscribe AMF: %s", err)

	} else if httpResponse.StatusCode != http.StatusCreated {
		logger.ConsumerLog.Errorf("Error on subscribe AMF: Excpect Status [%d] Received Status[%d]", http.StatusCreated, httpResponse.StatusCode)
	}
}
