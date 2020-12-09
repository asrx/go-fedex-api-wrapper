package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type UploadDocumentReferenceDetail struct {
	LineNumber *uint `xml:"LineNumber,omitempty"`

	CustomerReference string `xml:"CustomerReference,omitempty"`

	// Description of the uploaded document.
	Description string `xml:"Description,omitempty"`

	DocumentProducer *SimpleType.UploadDocumentProducerType `xml:"DocumentProducer,omitempty"`

	DocumentType *SimpleType.UploadDocumentType `xml:"DocumentType,omitempty"`

	DocumentId string `xml:"DocumentId,omitempty"`

	DocumentIdProducer *SimpleType.UploadDocumentIdProducer `xml:"DocumentIdProducer,omitempty"`
}
