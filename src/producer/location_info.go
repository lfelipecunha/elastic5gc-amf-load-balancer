package producer

import (
	"amfLoadBalancer/lib/http_wrapper"
	"amfLoadBalancer/lib/openapi/models"
	"amfLoadBalancer/src/context"
	"amfLoadBalancer/src/logger"
	"net/http"
)

func HandleProvideLocationInfoRequest(request *http_wrapper.Request) *http_wrapper.Response {
	logger.ProducerLog.Info("Handle Provide Location Info Request")

	requestLocInfo := request.Body.(models.RequestLocInfo)
	ueContextID := request.Params["ueContextId"]

	provideLocInfo, problemDetails := ProvideLocationInfoProcedure(requestLocInfo, ueContextID)
	if problemDetails != nil {
		return http_wrapper.NewResponse(int(problemDetails.Status), nil, problemDetails)
	} else {
		return http_wrapper.NewResponse(http.StatusOK, nil, provideLocInfo)
	}
}

func ProvideLocationInfoProcedure(requestLocInfo models.RequestLocInfo, ueContextID string) (
	*models.ProvideLocInfo, *models.ProblemDetails) {
	amfSelf := context.AMF_Self()

	ue, ok := amfSelf.AmfUeFindByUeContextID(ueContextID)
	if !ok {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
		return nil, problemDetails
	}

	anType := ue.GetAnType()
	if anType == "" {
		problemDetails := &models.ProblemDetails{
			Status: http.StatusNotFound,
			Cause:  "CONTEXT_NOT_FOUND",
		}
		return nil, problemDetails
	}

	provideLocInfo := new(models.ProvideLocInfo)

	ranUe := ue.RanUe[anType]
	if requestLocInfo.Req5gsLoc || requestLocInfo.ReqCurrentLoc {
		provideLocInfo.CurrentLoc = true
		provideLocInfo.Location = &ue.Location
	}

	if requestLocInfo.ReqRatType {
		provideLocInfo.RatType = ue.RatType
	}

	if requestLocInfo.ReqTimeZone {
		provideLocInfo.Timezone = ue.TimeZone
	}

	if requestLocInfo.SupportedFeatures != "" {
		provideLocInfo.SupportedFeatures = ranUe.SupportedFeatures
	}
	return provideLocInfo, nil
}
