package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackChargeDetail struct {
	Type *SimpleType.TrackChargeDetailType `xml:"Type,omitempty"`

	ChargeAmount *ComplexType.Money `xml:"ChargeAmount,omitempty"`
}
