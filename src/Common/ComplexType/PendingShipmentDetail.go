package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type PendingShipmentDetail struct {
	Type *SimpleType.PendingShipmentType `xml:"Type,omitempty"`

	// Date after which the pending shipment will no longer be available for completion.
	ExpirationDate string `xml:"ExpirationDate,omitempty"`

	ProcessingOptions *PendingShipmentProcessingOptionsRequested `xml:"ProcessingOptions,omitempty"`

	// These are documents that are recommended to be included with the shipment.
	RecommendedDocumentSpecification *RecommendedDocumentSpecification `xml:"RecommendedDocumentSpecification,omitempty"`
}
