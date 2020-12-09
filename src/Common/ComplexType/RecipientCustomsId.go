package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type RecipientCustomsId struct {

	// Specifies the kind of identification being used.
	Type *SimpleType.RecipientCustomsIdType `xml:"Type,omitempty"`

	// Contains the actual ID value, of the type specified above.
	Value string `xml:"Value,omitempty"`
}
