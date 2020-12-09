package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DocTabContentBarcoded struct {
	Symbology *SimpleType.BarcodeSymbologyType `xml:"Symbology,omitempty"`

	Specification *DocTabZoneSpecification `xml:"Specification,omitempty"`
}
