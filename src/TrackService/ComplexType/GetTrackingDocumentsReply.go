package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type GetTrackingDocumentsReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 GetTrackingDocumentsReply"`

	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	Documents []*TrackingDocument `xml:"Documents,omitempty"`
}
