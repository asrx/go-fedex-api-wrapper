package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RateDiscount struct {
	RateDiscountType *SimpleType.RateDiscountType `xml:"RateDiscountType,omitempty"`

	Description string `xml:"Description,omitempty"`

	Amount *Money `xml:"Amount,omitempty"`

	Percent float64 `xml:"Percent,omitempty"`
}
