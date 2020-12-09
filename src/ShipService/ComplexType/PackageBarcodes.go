package ComplexType

type PackageBarcodes struct {

	// Binary-style barcodes for this package.
	BinaryBarcodes []*BinaryBarcode `xml:"BinaryBarcodes,omitempty"`

	// String-style barcodes for this package.
	StringBarcodes []*StringBarcode `xml:"StringBarcodes,omitempty"`
}
