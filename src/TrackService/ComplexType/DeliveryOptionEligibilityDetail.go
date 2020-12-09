package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type DeliveryOptionEligibilityDetail struct {

	// Type of delivery option.
	Option *SimpleType.DeliveryOptionType `xml:"Option,omitempty"`

	// Eligibility of the customer for the specific delivery option.
	Eligibility *SimpleType.EligibilityType `xml:"Eligibility,omitempty"`
}
