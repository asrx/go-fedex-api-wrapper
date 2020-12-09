package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShippingDocumentEMailRecipient struct {

	// Identifies the relationship of this recipient in the shipment.
	RecipientType *SimpleType.EMailNotificationRecipientType `xml:"RecipientType,omitempty"`

	// Address to which the document is to be sent.
	Address string `xml:"Address,omitempty"`
}
