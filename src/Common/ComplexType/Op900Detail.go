package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Op900Detail struct {

	// Specifies characteristics of a shipping document to be produced.
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Identifies which reference type (from the package's customer references) is to be used as the source for the reference on this OP-900.
	Reference *SimpleType.CustomerReferenceType `xml:"Reference,omitempty"`

	// Specifies the usage and identification of customer supplied images to be used on this document.
	CustomerImageUsages []*CustomerImageUsage `xml:"CustomerImageUsages,omitempty"`

	// Data field to be used when a name is to be printed in the document instead of (or in addition to) a signature image.
	SignatureName string `xml:"SignatureName,omitempty"`
}
