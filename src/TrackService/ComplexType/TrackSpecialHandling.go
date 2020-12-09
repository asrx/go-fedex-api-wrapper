package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackSpecialHandling struct {
	Type *SimpleType.TrackSpecialHandlingType `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`

	PaymentType *SimpleType.TrackPaymentType `xml:"PaymentType,omitempty"`
}
