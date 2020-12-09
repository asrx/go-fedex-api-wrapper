package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type PackageRating struct {

	// This rate type identifies which entry in the following array is considered as presenting the "actual" rates for the package.
	ActualRateType *ReturnedRateType `xml:"ActualRateType,omitempty"`

	// The "list" net charge minus "actual" net charge.
	EffectiveNetDiscount *Money `xml:"EffectiveNetDiscount,omitempty"`

	// Each element of this field provides package-level rate data for a specific rate type.
	PackageRateDetails []*PackageRateDetail `xml:"PackageRateDetails,omitempty"`
}
