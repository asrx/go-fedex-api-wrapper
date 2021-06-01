package ComplexType

import (
	"encoding/xml"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
)

type DeleteTagRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/ship/v23 DeleteTagRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	// Only used for tags which had FedEx Express services.
	DispatchLocationId string `xml:"DispatchLocationId,omitempty"`

	// Only used for tags which had FedEx Express services.
	DispatchDate string `xml:"DispatchDate,omitempty"`

	// If the original ProcessTagRequest specified third-party payment, then the delete request must contain the same pay type and payor account number for security purposes.
	Payment *Payment `xml:"Payment,omitempty"`

	ConfirmationNumber string `xml:"ConfirmationNumber,omitempty"`
}
