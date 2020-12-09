package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type RecommendedDocumentSpecification struct {
	Types []*SimpleType.RecommendedDocumentType `xml:"Types,omitempty"`
}
