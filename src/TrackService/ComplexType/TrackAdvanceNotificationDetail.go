package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackAdvanceNotificationDetail struct {
	EstimatedTimeOfArrival string `xml:"EstimatedTimeOfArrival,omitempty"`

	Reason string `xml:"Reason,omitempty"`

	Status *SimpleType.TrackAdvanceNotificationStatusType `xml:"Status,omitempty"`

	StatusDescription string `xml:"StatusDescription,omitempty"`

	StatusTime string `xml:"StatusTime,omitempty"`
}
