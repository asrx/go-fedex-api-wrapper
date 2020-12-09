package ComplexType

type CommercialInvoiceDetail struct {
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Specifies the usage and identification of a customer supplied image to be used on this document.
	CustomerImageUsages []*CustomerImageUsage `xml:"CustomerImageUsages,omitempty"`
}
