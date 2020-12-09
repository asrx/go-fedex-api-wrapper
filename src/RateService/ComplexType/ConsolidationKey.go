package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/RateService/SimpleType"
)

type ConsolidationKey struct {

	// Specifies the type of consolidation.
	Type *SimpleType.ConsolidationType `xml:"Type,omitempty"`

	// Uniquely identifies the consolidation, within a given type and date.
	Index string `xml:"Index,omitempty"`

	// The date on which the consolidation was created.
	Date string `xml:"Date,omitempty"`
}
