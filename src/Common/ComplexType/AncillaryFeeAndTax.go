package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type AncillaryFeeAndTax struct {
	Type *SimpleType.AncillaryFeeAndTaxType `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`

	Amount *Money `xml:"Amount,omitempty"`
}
