package ComplexType


type CertificateOfOriginDetail struct {

	// Specifies characteristics of a shipping document to be produced.
	DocumentFormat *ShippingDocumentFormat `xml:"DocumentFormat,omitempty"`

	// Specifies the usage and identification of customer supplied images to be used on this document.
	CustomerImageUsages []*CustomerImageUsage `xml:"CustomerImageUsages,omitempty"`
}
