package handlers

import (
	"errors"
	"github.com/dhillondeep/go-uw-api"
	"google.golang.org/genproto/googleapis/cloud/dialogflow/v2"
	"strings"
)

func HandleRequest(req *dialogflow.WebhookRequest, uwClient *uwapi.UWAPI) (*dialogflow.WebhookResponse, error) {
	intentCat := strings.Split(req.QueryResult.Intent.DisplayName, "_")[1]

	switch intentCat {
	case CourseIntent:
		return HandleCourseReq(req.QueryResult, uwClient)
	default:
		return nil, errors.New("handler does not exist for intent category: " + intentCat)
	}
}