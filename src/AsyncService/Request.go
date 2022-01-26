package AsyncService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/AsyncService/ComplexType"
	//"github.com/asrx/go-fedex-api-wrapper/src/soap"
	"github.com/hooklift/gowsdl/soap"
)

type ASYNCPortType interface {
	RetrieveJobResults(request *RetrieveJobResultsRequest) (*RetrieveJobResultsReply, error)
}

type aSYNCPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/track"
var _TestUrl = "https://wsbeta.fedex.com:443/web-services/track"

func NewASYNCPortType(testEnv bool) ASYNCPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}

	return &aSYNCPortType{
		client: client,
	}
}

func (service *aSYNCPortType) RetrieveJobResults(request *RetrieveJobResultsRequest) (*RetrieveJobResultsReply, error) {
	response := new(RetrieveJobResultsReply)
	err := service.client.Call("http://fedex.com/ws/async/v4/retrieveJobResults", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
