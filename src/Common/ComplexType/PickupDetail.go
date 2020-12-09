package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type PickupDetail struct {
	ReadyDateTime string `xml:"ReadyDateTime,omitempty"`

	LatestPickupDateTime string `xml:"LatestPickupDateTime,omitempty"`

	CourierInstructions string `xml:"CourierInstructions,omitempty"`

	RequestType *SimpleType.PickupRequestType `xml:"RequestType,omitempty"`

	RequestSource *SimpleType.PickupRequestSourceType `xml:"RequestSource,omitempty"`
}
