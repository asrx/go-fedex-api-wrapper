package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ShipmentRating struct {

	// This rate type identifies which entry in the following array is considered as presenting the "actual" rates for the shipment.
	ActualRateType *SimpleType.ReturnedRateType `xml:"ActualRateType,omitempty"`

	// The "list" total net charge minus "actual" total net charge.
	EffectiveNetDiscount *ComplexType.Money `xml:"EffectiveNetDiscount,omitempty"`

	// Each element of this field provides shipment-level rate totals for a specific rate type.
	ShipmentRateDetails []*ComplexType.ShipmentRateDetail `xml:"ShipmentRateDetails,omitempty"`
}
