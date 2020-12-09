package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type HomeDeliveryPremiumDetail struct {
	HomeDeliveryPremiumType *SimpleType.HomeDeliveryPremiumType `xml:"HomeDeliveryPremiumType,omitempty"`

	Date string `xml:"Date,omitempty"`

	PhoneNumber string `xml:"PhoneNumber,omitempty"`
}
