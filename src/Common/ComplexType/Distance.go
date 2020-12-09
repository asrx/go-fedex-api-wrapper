package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Distance struct {

	// Identifies the distance quantity.
	Value float64 `xml:"Value,omitempty"`

	// Identifies the unit of measure for the distance value.
	Units *SimpleType.DistanceUnits `xml:"Units,omitempty"`
}
