package ComplexType

import . "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type TrackingDocumentEmailDetail struct {

	// Specifies the recipients of the email.
	Recipients []*EMailDetail `xml:"Recipients,omitempty"`

	// Identifies the person initiating the email.
	Sender *EMailDetail `xml:"Sender,omitempty"`

	// This is the localization of the email content.
	Localization *Localization `xml:"Localization,omitempty"`

	// A message included in the body of the email.
	PersonalMessage string `xml:"PersonalMessage,omitempty"`
}
