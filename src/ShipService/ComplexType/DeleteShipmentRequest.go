package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
	"encoding/xml"
	"time"
)

type DeleteShipmentRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/ship/v23 DeleteShipmentRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	ShipTimestamp time.Time `xml:"ShipTimestamp,omitempty"`

	TrackingId *TrackingId `xml:"TrackingId,omitempty"`

	DeletionControl *SimpleType.DeletionControlType `xml:"DeletionControl,omitempty"`
}
