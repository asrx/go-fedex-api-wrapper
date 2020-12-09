package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type GetTrackingDocumentsRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 GetTrackingDocumentsRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	SelectionDetails []*TrackSelectionDetail `xml:"SelectionDetails,omitempty"`

	TrackingDocumentSpecification *TrackingDocumentSpecification `xml:"TrackingDocumentSpecification,omitempty"`
}
