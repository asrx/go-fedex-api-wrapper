package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomerReference struct {
	CustomerReferenceType *SimpleType.CustomerReferenceType `xml:"CustomerReferenceType,omitempty"`

	Value string `xml:"Value,omitempty"`
}
