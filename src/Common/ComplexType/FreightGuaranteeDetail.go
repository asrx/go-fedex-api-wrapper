package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type FreightGuaranteeDetail struct {
	Type *SimpleType.FreightGuaranteeType `xml:"Type,omitempty"`

	// Date for all Freight guarantee types.
	Date string `xml:"Date,omitempty"`
}
