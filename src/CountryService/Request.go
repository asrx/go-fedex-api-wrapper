package CountryService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/CountryService/ComplexType"
	"github.com/hooklift/gowsdl/soap"
)

type CountryPortType interface {
	ValidatePostal(request *ValidatePostalRequest) (*ValidatePostalReply, error)
}

type countryPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/cnty"
var _TestUrl = "https://wsbeta.fedex.com:443/web-services/cnty"

func NewCountryPortType(testEnv bool) CountryPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}
	return &countryPortType{
		client: client,
	}
}

func (service *countryPortType) ValidatePostal(request *ValidatePostalRequest) (*ValidatePostalReply, error) {
	response := new(ValidatePostalReply)
	err := service.client.Call("http://fedex.com/ws/cnty/v8/validatePostal", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
