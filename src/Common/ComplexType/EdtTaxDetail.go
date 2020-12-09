package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type EdtTaxDetail struct {
	TaxType *SimpleType.EdtTaxType `xml:"TaxType,omitempty"`

	EffectiveDate string `xml:"EffectiveDate,omitempty"`

	Name string `xml:"Name,omitempty"`

	TaxableValue *Money `xml:"TaxableValue,omitempty"`

	Description string `xml:"Description,omitempty"`

	Formula string `xml:"Formula,omitempty"`

	Amount *Money `xml:"Amount,omitempty"`
}
