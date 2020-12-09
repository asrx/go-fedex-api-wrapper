package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShipmentEventNotificationDetail struct {
	AggregationType *SimpleType.ShipmentNotificationAggregationType `xml:"AggregationType,omitempty"`

	PersonalMessage string `xml:"PersonalMessage,omitempty"`

	EventNotifications []*ShipmentEventNotificationSpecification `xml:"EventNotifications,omitempty"`
}
