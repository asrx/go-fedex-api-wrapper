package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ShipmentRateDetail struct {

	// Type used for this specific set of rate data.
	RateType *SimpleType.ReturnedRateType `xml:"RateType,omitempty"`

	// Indicates the rate scale used.
	RateScale string `xml:"RateScale,omitempty"`

	// Indicates the rate zone used (based on origin and destination).
	RateZone string `xml:"RateZone,omitempty"`

	PricingCode *SimpleType.PricingCodeType `xml:"PricingCode,omitempty"`

	// Indicates which weight was used.
	RatedWeightMethod *SimpleType.RatedWeightMethod `xml:"RatedWeightMethod,omitempty"`

	// INTERNAL FEDEX USE ONLY.
	MinimumChargeType *SimpleType.MinimumChargeType `xml:"MinimumChargeType,omitempty"`

	// Specifies the currency exchange performed on financial amounts for this rate.
	CurrencyExchangeRate *CurrencyExchangeRate `xml:"CurrencyExchangeRate,omitempty"`

	// Indicates which special rating cases applied to this shipment.
	SpecialRatingApplied []*SimpleType.SpecialRatingAppliedType `xml:"SpecialRatingApplied,omitempty"`

	DimDivisor *uint `xml:"DimDivisor,omitempty"`

	// Identifies the type of dim divisor that was applied.
	DimDivisorType *SimpleType.RateDimensionalDivisorType `xml:"DimDivisorType,omitempty"`

	FuelSurchargePercent float64 `xml:"FuelSurchargePercent,omitempty"`

	TotalBillingWeight *Weight `xml:"TotalBillingWeight,omitempty"`

	// Sum of dimensional weights for all packages.
	TotalDimWeight *Weight `xml:"TotalDimWeight,omitempty"`

	TotalBaseCharge *Money `xml:"TotalBaseCharge,omitempty"`

	TotalFreightDiscounts *Money `xml:"TotalFreightDiscounts,omitempty"`

	TotalNetFreight *Money `xml:"TotalNetFreight,omitempty"`

	TotalSurcharges *Money `xml:"TotalSurcharges,omitempty"`

	// This shipment's totalNetFreight + totalSurcharges (not including totalTaxes).
	TotalNetFedExCharge *Money `xml:"TotalNetFedExCharge,omitempty"`

	// Total of the transportation-based taxes.
	TotalTaxes *Money `xml:"TotalTaxes,omitempty"`

	TotalNetCharge *Money `xml:"TotalNetCharge,omitempty"`

	TotalRebates *Money `xml:"TotalRebates,omitempty"`

	// Total of all values under this shipment's dutiesAndTaxes; only provided if estimated duties and taxes were calculated for this shipment.
	TotalDutiesAndTaxes *Money `xml:"TotalDutiesAndTaxes,omitempty"`

	// Identifies the total amount of the shipment-level fees and taxes that are not based on transportation charges or commodity-level estimated duties and taxes.
	TotalAncillaryFeesAndTaxes *Money `xml:"TotalAncillaryFeesAndTaxes,omitempty"`

	// The total of the totalDutiesAndTaxes plus the totalAncillaryFeesAndTaxes.
	TotalDutiesTaxesAndFees *Money `xml:"TotalDutiesTaxesAndFees,omitempty"`

	// This shipment's totalNetCharge + totalDutiesTaxesAndFees; some duties and taxes are only provided if estimated duties and taxes were calculated for this shipment AND duties, taxes and transportation charges are all paid by the same sender's account.
	TotalNetChargeWithDutiesAndTaxes *Money `xml:"TotalNetChargeWithDutiesAndTaxes,omitempty"`

	// Identifies the Rate Details per each leg in a Freight Shipment
	ShipmentLegRateDetails []*ShipmentLegRateDetail `xml:"ShipmentLegRateDetails,omitempty"`

	// Rate data specific to FedEx Freight and FedEx National Freight services.
	FreightRateDetail *FreightRateDetail `xml:"FreightRateDetail,omitempty"`

	// All rate discounts that apply to this shipment.
	FreightDiscounts []*RateDiscount `xml:"FreightDiscounts,omitempty"`

	// All rebates that apply to this shipment.
	Rebates []*Rebate `xml:"Rebates,omitempty"`

	// All surcharges that apply to this shipment.
	Surcharges []*Surcharge `xml:"Surcharges,omitempty"`

	// All transportation-based taxes applicable to this shipment.
	Taxes []*Tax `xml:"Taxes,omitempty"`

	// All commodity-based duties and taxes applicable to this shipment.
	DutiesAndTaxes []*EdtCommodityTax `xml:"DutiesAndTaxes,omitempty"`

	// Identifies the shipment-level fees and taxes that are not based on transportation charges or commodity-level estimated duties and taxes.
	AncillaryFeesAndTaxes []*AncillaryFeeAndTax `xml:"AncillaryFeesAndTaxes,omitempty"`

	// The "order level" variable handling charges.
	VariableHandlingCharges *VariableHandlingCharges `xml:"VariableHandlingCharges,omitempty"`

	// The total of all variable handling charges at both shipment (order) and package level.
	TotalVariableHandlingCharges *VariableHandlingCharges `xml:"TotalVariableHandlingCharges,omitempty"`
}
