package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type VariableHandlingChargeDetail struct {
	FixedValue *Money `xml:"FixedValue,omitempty"`

	// Actual percentage (10 means 10%, which is a mutiplier of 0.1)
	PercentValue float64 `xml:"PercentValue,omitempty"`

	// Select the value from a set of rate data to which the percentage is applied.
	RateElementBasis SimpleType.RateElementBasisType `xml:"RateElementBasis,omitempty"`

	// Select the type of rate from which the element is to be selected.
	RateTypeBasis SimpleType.RateTypeBasisType `xml:"RateTypeBasis,omitempty"`
}
