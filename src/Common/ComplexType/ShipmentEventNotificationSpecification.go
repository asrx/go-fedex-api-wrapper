package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShipmentEventNotificationSpecification struct {
	Role *SimpleType.ShipmentNotificationRoleType `xml:"Role,omitempty"`

	Events []*SimpleType.NotificationEventType `xml:"Events,omitempty"`

	NotificationDetail *NotificationDetail `xml:"NotificationDetail,omitempty"`

	FormatSpecification *ShipmentNotificationFormatSpecification `xml:"FormatSpecification,omitempty"`
}
