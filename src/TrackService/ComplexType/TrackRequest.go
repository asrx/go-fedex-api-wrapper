package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
	"encoding/xml"
)

type TrackRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 TrackRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	// Descriptive data identifying the client submitting the transaction.
	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	// Contains a free form field that is echoed back in the reply to match requests with replies and data that governs the data payload language/translations.
	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	// The version of the request being used.
	Version *VersionId `xml:"Version,omitempty"`

	// Specifies the details needed to select the shipment being requested to be tracked.
	SelectionDetails []*TrackSelectionDetail `xml:"SelectionDetails,omitempty"`

	// The customer can specify a desired time out value for this particular transaction.
	TransactionTimeOutValueInMilliseconds *uint `xml:"TransactionTimeOutValueInMilliseconds,omitempty"`

	ProcessingOptions []*TrackRequestProcessingOptionType `xml:"ProcessingOptions,omitempty"`
}
