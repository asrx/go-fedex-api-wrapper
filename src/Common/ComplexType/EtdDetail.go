package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type EtdDetail struct {
	Attributes []*SimpleType.EtdAttributeType `xml:"Attributes,omitempty"`

	// Indicates the types of shipping documents produced for the shipper by FedEx (see ShippingDocumentSpecification) which should be copied back to the shipper in the shipment result data.
	RequestedDocumentCopies []*SimpleType.RequestedShippingDocumentType `xml:"RequestedDocumentCopies,omitempty"`

	DocumentReferences []*UploadDocumentReferenceDetail `xml:"DocumentReferences,omitempty"`
}
