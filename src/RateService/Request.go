package RateService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/RateService/ComplexType"
	"github.com/hooklift/gowsdl/soap"
)

type RatePortType interface {
	GetRates(request *RateRequest) (*RateReply, error)
}

type ratePortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/rate"
var _TestUrl = "https://wsbeta.fedex.com:443/web-services/rate"

func NewRatePortType(testEnv bool) RatePortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}
	return &ratePortType{
		client: client,
	}
}

func (service *ratePortType) GetRates(request *RateRequest) (*RateReply, error) {
	response := new(RateReply)
	err := service.client.Call("http://fedex.com/ws/rate/v24/getRates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}