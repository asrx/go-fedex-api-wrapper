package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type SendNotificationsRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 SendNotificationsRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	// Descriptive data identifying the client submitting the transaction.
	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	// Contains a free form field that is echoed back in the reply to match requests with replies and data that governs the data payload language/translations
	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	// The tracking number to which the notifications will be triggered from.
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	// Indicates whether to return tracking information for all associated packages.
	MultiPiece bool `xml:"MultiPiece,omitempty"`

	// When the MoreDataAvailable field is true in a TrackNotificationReply the PagingToken must be sent in the subsequent TrackNotificationRequest to retrieve the next page of data.
	PagingToken string `xml:"PagingToken,omitempty"`

	// Use this field when your original request informs you that there are duplicates of this tracking number. If you get duplicates you will also receive some information about each of the duplicate tracking numbers to enable you to chose one and resend that number along with the TrackingNumberUniqueId to get notifications for that tracking number.
	TrackingNumberUniqueId string `xml:"TrackingNumberUniqueId,omitempty"`

	// To narrow the search to a period in time the ShipDateRangeBegin and ShipDateRangeEnd can be used to help eliminate duplicates.
	ShipDateRangeBegin string `xml:"ShipDateRangeBegin,omitempty"`

	// To narrow the search to a period in time the ShipDateRangeBegin and ShipDateRangeEnd can be used to help eliminate duplicates.
	ShipDateRangeEnd string `xml:"ShipDateRangeEnd,omitempty"`

	// Included in the email notification identifying the requester of this notification.
	SenderEMailAddress string `xml:"SenderEMailAddress,omitempty"`

	// Included in the email notification identifying the requester of this notification.
	SenderContactName string `xml:"SenderContactName,omitempty"`

	// This replaces eMailNotificationDetail
	EventNotificationDetail *ShipmentEventNotificationDetail `xml:"EventNotificationDetail,omitempty"`
}
