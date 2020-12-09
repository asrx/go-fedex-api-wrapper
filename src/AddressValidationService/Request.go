package AddressValidationService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/AddressValidationService/ComplexType"
	"github.com/hooklift/gowsdl/soap"
)

type AddressValidationPortType interface {
	AddressValidation(request *AddressValidationRequest) (*AddressValidationReply, error)
}


type addressValidationPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/addressvalidation"
var _TestUrl = "https://wsbeta.fedex.com:443/web-services/addressvalidation"

func NewAddressValidationPortType(testEnv bool) AddressValidationPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}

	return &addressValidationPortType{
		client: client,
	}
}

func (service *addressValidationPortType) AddressValidation(request *AddressValidationRequest) (*AddressValidationReply, error) {
	response := new(AddressValidationReply)
	err := service.client.Call("http://fedex.com/ws/addressvalidation/v4/addressValidation", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
