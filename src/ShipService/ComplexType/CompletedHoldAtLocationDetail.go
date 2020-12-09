package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CompletedHoldAtLocationDetail struct {

	// Identifies the branded location name, the hold at location phone number and the address of the location.
	HoldingLocation *ComplexType.ContactAndAddress `xml:"HoldingLocation,omitempty"`

	// Identifies the type of FedEx location.
	HoldingLocationType *SimpleType.FedExLocationType `xml:"HoldingLocationType,omitempty"`

	HoldingLocationTypeForDisplay string `xml:"HoldingLocationTypeForDisplay,omitempty"`
}
