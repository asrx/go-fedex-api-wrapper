package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ShippingDocumentEMailDetail struct {

	// Provides the roles and email addresses for e-mail recipients.
	EMailRecipients []*ShippingDocumentEMailRecipient `xml:"EMailRecipients,omitempty"`

	// Identifies the convention by which documents are to be grouped as e-mail attachments.
	Grouping *SimpleType.ShippingDocumentEMailGroupingType `xml:"Grouping,omitempty"`

	// Specifies the language in which the email containing the document is requested to be composed.
	Localization *Localization `xml:"Localization,omitempty"`
}
