package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type AdditionalLabelsDetail struct {

	// The type of additional labels to return.
	Type *SimpleType.AdditionalLabelsType `xml:"Type,omitempty"`

	// The number of this type label to return
	Count *uint `xml:"Count,omitempty"`
}

