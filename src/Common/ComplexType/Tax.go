package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type Tax struct {
	TaxType *SimpleType.TaxType `xml:"TaxType,omitempty"`

	Description string `xml:"Description,omitempty"`

	Amount *Money `xml:"Amount,omitempty"`
}
