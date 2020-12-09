package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type Surcharge struct {
	SurchargeType *SimpleType.SurchargeType `xml:"SurchargeType,omitempty"`

	Level *SimpleType.SurchargeLevelType `xml:"Level,omitempty"`

	Description string `xml:"Description,omitempty"`

	Amount *Money `xml:"Amount,omitempty"`
}
