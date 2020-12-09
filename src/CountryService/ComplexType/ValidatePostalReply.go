package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type ValidatePostalReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/cnty/v8 ValidatePostalReply"`

	HighestSeverity *SimpleType.NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*ComplexType.Notification `xml:"Notifications,omitempty"`

	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	PostalDetail *PostalDetail `xml:"PostalDetail,omitempty"`
}
