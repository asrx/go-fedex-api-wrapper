package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type RegulatoryLabelContentDetail struct {
	Type *SimpleType.RegulatoryLabelType `xml:"Type,omitempty"`

	// Specifies how the customer requested the regulatory label to be generated.
	GenerationOptions []*SimpleType.CustomerSpecifiedLabelGenerationOptionType `xml:"GenerationOptions,omitempty"`
}
