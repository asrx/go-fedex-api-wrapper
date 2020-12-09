package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackPayment struct {

	// Indicates the classification of the charges being paid.
	Classification *SimpleType.TrackChargesPaymentClassificationType `xml:"Classification,omitempty"`

	Type *SimpleType.TrackPaymentType `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`
}
