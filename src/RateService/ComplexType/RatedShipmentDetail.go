package ComplexType

import . "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type RatedShipmentDetail struct {

	// The difference between "list" and "account" total net charge.
	EffectiveNetDiscount *Money `xml:"EffectiveNetDiscount,omitempty"`

	// Express COD is shipment level.
	AdjustedCodCollectionAmount *Money `xml:"AdjustedCodCollectionAmount,omitempty"`

	// The shipment-level totals for this rate type.
	ShipmentRateDetail *ShipmentRateDetail `xml:"ShipmentRateDetail,omitempty"`

	// The package-level data for this rate type.
	RatedPackages []*RatedPackageDetail `xml:"RatedPackages,omitempty"`
}
