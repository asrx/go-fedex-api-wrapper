package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type PrintedReference struct {
	Type *SimpleType.PrintedReferenceType `xml:"Type,omitempty"`

	Value string `xml:"Value,omitempty"`
}
