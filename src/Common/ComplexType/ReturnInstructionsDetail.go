package ComplexType

type ReturnInstructionsDetail struct {
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Specifies additional customer provided text to be inserted into the return document.
	CustomText string `xml:"CustomText,omitempty"`
}
