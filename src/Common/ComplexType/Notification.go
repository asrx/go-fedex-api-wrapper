package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Notification struct {

	// The severity of this notification. This can indicate success or failure or some other information about the request. The values that can be returned are SUCCESS - Your transaction succeeded with no other applicable information. NOTE - Additional information that may be of interest to you about your transaction. WARNING - Additional information that you need to know about your transaction that you may need to take action on. ERROR - Information about an error that occurred while processing your transaction. FAILURE - FedEx was unable to process your transaction at this time due to a system failure. Please try again later
	Severity *SimpleType.NotificationSeverityType `xml:"Severity,omitempty"`

	// Indicates the source of this notification. Combined with the Code it uniquely identifies this notification
	Source string `xml:"Source,omitempty"`

	// A code that represents this notification. Combined with the Source it uniquely identifies this notification.
	Code string `xml:"Code,omitempty"`

	// Human-readable text that explains this notification.
	Message string `xml:"Message,omitempty"`

	// The translated message. The language and locale specified in the ClientDetail. Localization are used to determine the representation. Currently only supported in a TrackReply.
	LocalizedMessage string `xml:"LocalizedMessage,omitempty"`

	// A collection of name/value pairs that provide specific data to help the client determine the nature of an error (or warning, etc.) without having to parse the message string.
	MessageParameters []*NotificationParameter `xml:"MessageParameters,omitempty"`
}
