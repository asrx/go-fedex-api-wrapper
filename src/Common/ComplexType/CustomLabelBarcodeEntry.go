package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomLabelBarcodeEntry struct {
	Position *CustomLabelPosition `xml:"Position,omitempty"`

	Format string `xml:"Format,omitempty"`

	DataFields []string `xml:"DataFields,omitempty"`

	BarHeight int32 `xml:"BarHeight,omitempty"`

	// Width of thinnest bar/space element in the barcode.
	ThinBarWidth int32 `xml:"ThinBarWidth,omitempty"`

	BarcodeSymbology *SimpleType.BarcodeSymbologyType `xml:"BarcodeSymbology,omitempty"`
}