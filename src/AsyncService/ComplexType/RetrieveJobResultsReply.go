package ComplexType

import (
	"encoding/xml"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RetrieveJobResultsReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/async/v4 RetrieveJobResultsReply"`

	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	Artifacts []*RetrievedArtifact `xml:"Artifacts,omitempty"`
}
