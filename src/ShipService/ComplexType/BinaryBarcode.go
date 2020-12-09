package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type BinaryBarcode struct {

	// The kind of barcode data in this instance.
	Type *SimpleType.BinaryBarcodeType `xml:"Type,omitempty"`

	// The data content of this instance.
	Value []byte `xml:"Value,omitempty"`
}

