package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type CompletedEtdDetail struct {

	// The identifier for all clearance documents associated with this shipment.
	FolderId string `xml:"FolderId,omitempty"`

	Type *SimpleType.CompletedEtdType `xml:"Type,omitempty"`

	UploadDocumentReferenceDetails []*ComplexType.UploadDocumentReferenceDetail `xml:"UploadDocumentReferenceDetails,omitempty"`
}
