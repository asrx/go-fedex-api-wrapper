package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
	"time"
)

type ValidatePostalRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/cnty/v8 ValidatePostalRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *ComplexType.WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ComplexType.ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	ShipDateTime time.Time `xml:"ShipDateTime,omitempty"`

	Address *ComplexType.Address `xml:"Address,omitempty"`

	RoutingCode string `xml:"RoutingCode,omitempty"`

	CheckForMismatch bool `xml:"CheckForMismatch,omitempty"`

	CarrierCode *SimpleType.CarrierCodeType `xml:"CarrierCode,omitempty"`
}
