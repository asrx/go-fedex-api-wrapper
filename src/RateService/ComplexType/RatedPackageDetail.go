package ComplexType

import (
	."github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RatedPackageDetail struct {

	// Echoed from the corresponding package in the rate request (if provided).
	TrackingIds []*TrackingId `xml:"TrackingIds,omitempty"`

	// Used with request containing PACKAGE_GROUPS, to identify which group of identical packages was used to produce a reply item.
	GroupNumber *uint `xml:"GroupNumber,omitempty"`

	// The difference between "list" and "account" net charge.
	EffectiveNetDiscount *Money `xml:"EffectiveNetDiscount,omitempty"`

	// Ground COD is shipment level.
	AdjustedCodCollectionAmount *Money `xml:"AdjustedCodCollectionAmount,omitempty"`

	OversizeClass *OversizeClassType `xml:"OversizeClass,omitempty"`

	// Rate data that are tied to a specific package and rate type combination.
	PackageRateDetail *PackageRateDetail `xml:"PackageRateDetail,omitempty"`
}
