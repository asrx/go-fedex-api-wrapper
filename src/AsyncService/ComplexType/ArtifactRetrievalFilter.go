package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/AsyncService/SimpleType"

type ArtifactRetrievalFilter struct {
	AccessReference string `xml:"AccessReference,omitempty"`

	Type *SimpleType.ArtifactType `xml:"Type,omitempty"`

	ReferenceId string `xml:"ReferenceId,omitempty"`
}
