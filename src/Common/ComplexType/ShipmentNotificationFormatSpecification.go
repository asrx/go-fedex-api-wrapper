package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShipmentNotificationFormatSpecification struct {
	Type *SimpleType.NotificationFormatType `xml:"Type,omitempty"`
}