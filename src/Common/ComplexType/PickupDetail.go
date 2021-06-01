package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"time"
)

type PickupDetail struct {
	ReadyDateTime time.Time `xml:"ReadyDateTime,omitempty"`

	LatestPickupDateTime time.Time `xml:"LatestPickupDateTime,omitempty"`

	CourierInstructions string `xml:"CourierInstructions,omitempty"`

	RequestType *SimpleType.PickupRequestType `xml:"RequestType,omitempty"`

	RequestSource *SimpleType.PickupRequestSourceType `xml:"RequestSource,omitempty"`
}
