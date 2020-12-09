package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type ProcessShipmentReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/ship/v23 ProcessShipmentReply"`

	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	JobId string `xml:"JobId,omitempty"`

	CompletedShipmentDetail *CompletedShipmentDetail `xml:"CompletedShipmentDetail,omitempty"`

	// Empty unless error label behavior is PACKAGE_ERROR_LABELS and one or more errors occured during transaction processing.
	ErrorLabels []*ShippingDocument `xml:"ErrorLabels,omitempty"`
}
