package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type TrackReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 TrackReply"`

	// This contains the severity type of the most severe Notification in the Notifications array.
	HighestSeverity *SimpleType.NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	// Information about the request/reply such was the transaction successful or not, and any additional information relevant to the request and/or reply. There may be multiple Notifications in a reply.
	Notifications []*ComplexType.Notification `xml:"Notifications,omitempty"`

	// Contains the CustomerTransactionDetail that is echoed back to the caller for matching requests and replies and a Localization element for defining the language/translation used in the reply data.
	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	// Contains the version of the reply being used.
	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	// Contains detailed tracking entity information.
	CompletedTrackDetails []*CompletedTrackDetail `xml:"CompletedTrackDetails,omitempty"`
}
