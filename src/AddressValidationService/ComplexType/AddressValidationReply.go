package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type AddressValidationReply struct {

	XMLName xml.Name `xml:"http://fedex.com/ws/addressvalidation/v4 AddressValidationReply"`

	HighestSeverity *SimpleType.NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*SimpleType.NotificationSeverityType `xml:"Notifications,omitempty"`

	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	ReplyTimestamp string `xml:"ReplyTimestamp,omitempty"`

	AddressResults []*AddressValidationResult `xml:"AddressResults,omitempty"`
}
