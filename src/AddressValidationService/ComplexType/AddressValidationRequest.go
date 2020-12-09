package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"encoding/xml"
	"time"
)

type AddressValidationRequest struct {

	XMLName xml.Name `xml:"http://fedex.com/ws/addressvalidation/v4 AddressValidationRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *ComplexType.WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ComplexType.ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *ComplexType.TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *ComplexType.VersionId `xml:"Version,omitempty"`

	InEffectAsOfTimestamp time.Time `xml:"InEffectAsOfTimestamp,omitempty"`

	AddressesToValidate []*AddressToValidate `xml:"AddressesToValidate,omitempty"`
}
