package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type RadionuclideActivity struct {
	Value float64 `xml:"Value,omitempty"`

	UnitOfMeasure *SimpleType.RadioactivityUnitOfMeasure `xml:"UnitOfMeasure,omitempty"`
}
