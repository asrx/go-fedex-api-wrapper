package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackingDocumentSpecification struct {
	DocumentTypes []*SimpleType.TrackingDocumentType `xml:"DocumentTypes,omitempty"`

	BillOfLadingDocumentDetail *TrackingBillOfLadingDocumentDetail `xml:"BillOfLadingDocumentDetail,omitempty"`

	FreightBillingDocumentDetail *TrackingFreightBillingDocumentDetail `xml:"FreightBillingDocumentDetail,omitempty"`

	SignatureProofOfDeliveryDetail *TrackingSignatureProofOfDeliveryDetail `xml:"SignatureProofOfDeliveryDetail,omitempty"`
}
