package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type PackageRateDetail struct {

	// Type used for this specific set of rate data.
	RateType *SimpleType.ReturnedRateType `xml:"RateType,omitempty"`

	// Indicates which weight was used.
	RatedWeightMethod *SimpleType.RatedWeightMethod `xml:"RatedWeightMethod,omitempty"`

	// INTERNAL FEDEX USE ONLY.
	MinimumChargeType *SimpleType.MinimumChargeType `xml:"MinimumChargeType,omitempty"`

	BillingWeight *Weight `xml:"BillingWeight,omitempty"`

	// The dimensional weight of this package (if greater than actual).
	DimWeight *Weight `xml:"DimWeight,omitempty"`

	// The oversize weight of this package (if the package is oversize).
	OversizeWeight *Weight `xml:"OversizeWeight,omitempty"`

	// The transportation charge only (prior to any discounts applied) for this package.
	BaseCharge *Money `xml:"BaseCharge,omitempty"`

	// The sum of all discounts on this package.
	TotalFreightDiscounts *Money `xml:"TotalFreightDiscounts,omitempty"`

	// This package's baseCharge - totalFreightDiscounts.
	NetFreight *Money `xml:"NetFreight,omitempty"`

	// The sum of all surcharges on this package.
	TotalSurcharges *Money `xml:"TotalSurcharges,omitempty"`

	// This package's netFreight + totalSurcharges (not including totalTaxes).
	NetFedExCharge *Money `xml:"NetFedExCharge,omitempty"`

	// The sum of all taxes on this package.
	TotalTaxes *Money `xml:"TotalTaxes,omitempty"`

	// This package's netFreight + totalSurcharges + totalTaxes.
	NetCharge *Money `xml:"NetCharge,omitempty"`

	TotalRebates *Money `xml:"TotalRebates,omitempty"`

	// All rate discounts that apply to this package.
	FreightDiscounts []*RateDiscount `xml:"FreightDiscounts,omitempty"`

	// All rebates that apply to this package.
	Rebates []*Rebate `xml:"Rebates,omitempty"`

	// All surcharges that apply to this package (either because of characteristics of the package itself, or because it is carrying per-shipment surcharges for the shipment of which it is a part).
	Surcharges []*Surcharge `xml:"Surcharges,omitempty"`

	// All taxes applicable (or distributed to) this package.
	Taxes []*Tax `xml:"Taxes,omitempty"`

	VariableHandlingCharges *VariableHandlingCharges `xml:"VariableHandlingCharges,omitempty"`
}
