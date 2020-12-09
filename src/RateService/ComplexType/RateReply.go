package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type RateReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/rate/v24 RateReply"`

	// This indicates the highest level of severity of all the notifications returned in this reply.
	HighestSeverity *SimpleType2.NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	// The descriptive data regarding the results of the submitted transaction.
	Notifications []*ComplexType.Notification `xml:"Notifications,omitempty"`

	// Contains the CustomerTransactionId that was sent in the request.
	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	// The version of this reply.
	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	// Each element contains all rate data for a single service. If service was specified in the request, there will be a single entry in this array; if service was omitted in the request, there will be a separate entry in this array for each service being compared.
	RateReplyDetails []*RateReplyDetail `xml:"RateReplyDetails,omitempty"`
}
