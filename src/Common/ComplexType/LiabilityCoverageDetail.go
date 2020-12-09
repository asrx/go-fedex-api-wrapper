package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type LiabilityCoverageDetail struct {
	CoverageType *SimpleType.LiabilityCoverageType `xml:"CoverageType,omitempty"`

	// Identifies the Liability Coverage Amount. For Jan 2010 this value represents coverage amount per pound
	CoverageAmount *Money `xml:"CoverageAmount,omitempty"`
}
