package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type StringBarcode struct {

	// The kind of barcode data in this instance.
	Type *SimpleType.StringBarcodeType `xml:"Type,omitempty"`

	// The data content of this instance.
	Value string `xml:"Value,omitempty"`
}
