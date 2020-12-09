package TrackService

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/TrackService/ComplexType"
	"github.com/hooklift/gowsdl/soap"
)


type TrackPortType interface {
	Track(request *TrackRequest) (*TrackReply, error)

	GetTrackingDocuments(request *GetTrackingDocumentsRequest) (*GetTrackingDocumentsReply, error)

	SendNotifications(request *SendNotificationsRequest) (*SendNotificationsReply, error)
}

type trackPortType struct {
	client *soap.Client
}

var _ProductionUrl = "https://ws.fedex.com:443/web-services/track"
var _TestUrl = "https://ws.fedex.com:443/web-services/track"

func NewTrackPortType(testEnv bool) TrackPortType {
	var client = new(soap.Client)
	if testEnv {
		client = soap.NewClient(_TestUrl)
	}else {
		client = soap.NewClient(_ProductionUrl)
	}

	return &trackPortType{
		client: client,
	}
}

func (service *trackPortType) Track(request *TrackRequest) (*TrackReply, error) {
	response := new(TrackReply)
	err := service.client.Call("http://fedex.com/ws/track/v16/track", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *trackPortType) GetTrackingDocuments(request *GetTrackingDocumentsRequest) (*GetTrackingDocumentsReply, error) {
	response := new(GetTrackingDocumentsReply)
	err := service.client.Call("http://fedex.com/ws/track/v16/getTrackingDocuments", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *trackPortType) SendNotifications(request *SendNotificationsRequest) (*SendNotificationsReply, error) {
	response := new(SendNotificationsReply)
	err := service.client.Call("http://fedex.com/ws/track/v16/sendNotifications", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
