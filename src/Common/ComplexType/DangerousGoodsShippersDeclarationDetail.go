package ComplexType

type DangerousGoodsShippersDeclarationDetail struct {

	// Specifies characteristics of a shipping document to be produced.
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Specifies the usage and identification of customer supplied images to be used on this document.
	CustomerImageUsages []*CustomerImageUsage `xml:"CustomerImageUsages,omitempty"`
}
