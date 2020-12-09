package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type DocumentRequirementsDetail struct {
	RequiredDocuments []*SimpleType.RequiredDocumentType `xml:"RequiredDocuments,omitempty"`

	GenerationDetails []*DocumentGenerationDetail `xml:"GenerationDetails,omitempty"`

	// Lists the documents that are not accepted by FedEx for this shipment.
	ProhibitedDocuments []*SimpleType.EnterpriseDocumentType `xml:"ProhibitedDocuments,omitempty"`
}
