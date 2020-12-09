package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ShippingDocumentFormat struct {

	// Specifies how to create, organize, and return the document.
	Dispositions []*ShippingDocumentDispositionDetail `xml:"Dispositions,omitempty"`

	// Specifies how far down the page to move the beginning of the image; allows for printing on letterhead and other pre-printed stock.
	TopOfPageOffset *LinearMeasure `xml:"TopOfPageOffset,omitempty"`

	ImageType *SimpleType.ShippingDocumentImageType `xml:"ImageType,omitempty"`

	StockType *SimpleType.ShippingDocumentStockType `xml:"StockType,omitempty"`

	// For those shipping document types which have both a "form" and "instructions" component (e.g. NAFTA Certificate of Origin and General Agency Agreement), this field indicates whether to provide the instructions.
	ProvideInstructions bool `xml:"ProvideInstructions,omitempty"`

	OptionsRequested *DocumentFormatOptionsRequested `xml:"OptionsRequested,omitempty"`

	// Governs the language to be used for this individual document, independently from other content returned for the same shipment.
	Localization *Localization `xml:"Localization,omitempty"`
}
