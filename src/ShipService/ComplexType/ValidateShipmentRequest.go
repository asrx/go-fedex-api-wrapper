package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
)

type ValidateShipmentRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/ship/v23 ValidateShipmentRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	RequestedShipment *RequestedShipment `xml:"RequestedShipment,omitempty"`
}
