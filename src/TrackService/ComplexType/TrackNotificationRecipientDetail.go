package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type TrackNotificationRecipientDetail struct {
	NotificationEventsAvailable []*SimpleType.NotificationEventType `xml:"NotificationEventsAvailable,omitempty"`
}
