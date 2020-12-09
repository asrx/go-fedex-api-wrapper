package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomsOptionDetail struct {
	Type *SimpleType.CustomsOptionType `xml:"Type,omitempty"`

	// Specifies additional description about customs options. This is a required field when the customs options type is "OTHER".
	Description string `xml:"Description,omitempty"`
}
