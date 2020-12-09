package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomDocumentDetail struct {

	// Common information controlling document production.
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Applicable only to documents produced on thermal printers with roll stock.
	LabelPrintingOrientation *SimpleType.LabelPrintingOrientationType `xml:"LabelPrintingOrientation,omitempty"`

	// Applicable only to documents produced on thermal printers with roll stock.
	LabelRotation *SimpleType.LabelRotationType `xml:"LabelRotation,omitempty"`

	// Identifies the formatting specification used to construct this custom document.
	SpecificationId string `xml:"SpecificationId,omitempty"`

	CustomContent *CustomLabelDetail `xml:"CustomContent,omitempty"`
}
