package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)


type NotificationDetail struct {

	// Indicates the type of notification that will be sent.
	NotificationType *SimpleType.NotificationType `xml:"NotificationType,omitempty"`

	// Specifies the email notification details.
	EmailDetail *EMailDetail `xml:"EmailDetail,omitempty"`

	// Specifies the localization for this notification.
	Localization *Localization `xml:"Localization,omitempty"`
}
