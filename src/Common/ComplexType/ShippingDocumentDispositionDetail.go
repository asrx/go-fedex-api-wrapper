package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShippingDocumentDispositionDetail struct {

	// Values in this field specify how to create and return the document.
	DispositionType *SimpleType.ShippingDocumentDispositionType `xml:"DispositionType,omitempty"`

	// Specifies how to organize all documents of this type.
	Grouping *SimpleType.ShippingDocumentGroupingType `xml:"Grouping,omitempty"`

	// Specifies how to e-mail document images.
	EMailDetail *ShippingDocumentEMailDetail `xml:"EMailDetail,omitempty"`

	// Specifies how a queued document is to be printed.
	PrintDetail *ShippingDocumentPrintDetail `xml:"PrintDetail,omitempty"`
}
