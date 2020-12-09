package ComplexType

import . "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type FreightRateDetail struct {

	// A unique identifier for a specific rate quotation.
	QuoteNumber string `xml:"QuoteNumber,omitempty"`

	// Specifies whether the rate quote was automated or manual.
	QuoteType *FreightRateQuoteType `xml:"QuoteType,omitempty"`

	// Specifies how total base charge is determined.
	BaseChargeCalculation *FreightBaseChargeCalculationType `xml:"BaseChargeCalculation,omitempty"`

	// Freight charges which accumulate to the total base charge for the shipment.
	BaseCharges []*FreightBaseCharge `xml:"BaseCharges,omitempty"`

	// Human-readable descriptions of additional information on this shipment rating.
	Notations []*FreightRateNotation `xml:"Notations,omitempty"`
}
