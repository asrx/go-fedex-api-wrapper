package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type ShippingDocument struct {
	Type *ReturnedShippingDocumentType `xml:"Type,omitempty"`

	// The localizations are populated if multiple language versions of a shipping document are returned.
	Localizations []*Localization `xml:"Localizations,omitempty"`

	// Specifies how this document image/file is organized.
	Grouping *ShippingDocumentGroupingType `xml:"Grouping,omitempty"`

	ShippingDocumentDisposition *ShippingDocumentDispositionType `xml:"ShippingDocumentDisposition,omitempty"`

	// The name under which a STORED, DEFERRED or EMAILED document is written.
	AccessReference string `xml:"AccessReference,omitempty"`

	// Specifies the image type of this shipping document.
	ImageType *ShippingDocumentImageType `xml:"ImageType,omitempty"`

	// Specifies the image resolution in DPI (dots per inch).
	Resolution *uint `xml:"Resolution,omitempty"`

	// Can be zero for documents whose disposition implies that no content is included.
	CopiesToPrint *uint `xml:"CopiesToPrint,omitempty"`

	// One or more document parts which make up a single logical document, such as multiple pages of a single form.
	Parts []*ShippingDocumentPart `xml:"Parts,omitempty"`
}
