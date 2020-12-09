package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type ShipmentReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/ship/v23 ShipmentReply"`

	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`
}
