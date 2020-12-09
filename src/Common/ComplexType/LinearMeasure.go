package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type LinearMeasure struct {

	// The numerical quantity of this measurement.
	Value float64 `xml:"Value,omitempty"`

	// The units for this measurement.
	Units *SimpleType.LinearUnits `xml:"Units,omitempty"`
}
