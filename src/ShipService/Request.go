package ShipService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/ShipService/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/soap"

	//"github.com/hooklift/gowsdl/soap"
)

type ShipPortType interface {
	ProcessTag(request *ProcessTagRequest) (*ProcessTagReply, error)

	ProcessShipment(request *ProcessShipmentRequest) (*ProcessShipmentReply, error)

	DeleteTag(request *DeleteTagRequest) (*ShipmentReply, error)

	DeleteShipment(request *DeleteShipmentRequest) (*ShipmentReply, error)

	ValidateShipment(request *ValidateShipmentRequest) (*ShipmentReply, error)
}

type shipPortType struct {
	client *soap.Client
}

func NewShipPortType(testEnv bool) ShipPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}

	return &shipPortType{
		client: client,
	}
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/ship"
var _TestUrl = "https://wsbeta.fedex.com:443/web-services/ship"

func (service *shipPortType) ProcessTag(request *ProcessTagRequest) (*ProcessTagReply, error) {
	response := new(ProcessTagReply)
	err := service.client.Call("http://fedex.com/ws/ship/v23/processTag", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *shipPortType) ProcessShipment(request *ProcessShipmentRequest) (*ProcessShipmentReply, error) {
	response := new(ProcessShipmentReply)
	err := service.client.Call("http://fedex.com/ws/ship/v23/processShipment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *shipPortType) DeleteTag(request *DeleteTagRequest) (*ShipmentReply, error) {
	response := new(ShipmentReply)
	err := service.client.Call("http://fedex.com/ws/ship/v23/deleteTag", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *shipPortType) DeleteShipment(request *DeleteShipmentRequest) (*ShipmentReply, error) {
	response := new(ShipmentReply)
	err := service.client.Call("http://fedex.com/ws/ship/v23/deleteShipment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *shipPortType) ValidateShipment(request *ValidateShipmentRequest) (*ShipmentReply, error) {
	response := new(ShipmentReply)
	err := service.client.Call("http://fedex.com/ws/ship/v23/validateShipment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

