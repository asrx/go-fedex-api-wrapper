package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CommercialInvoice struct {

	// Any comments that need to be communicated about this shipment.
	Comments []string `xml:"Comments,omitempty"`

	// Any freight charges that are associated with this shipment.
	FreightCharge *Money `xml:"FreightCharge,omitempty"`

	// Any taxes or miscellaneous charges(other than Freight charges or Insurance charges) that are associated with this shipment.
	TaxesOrMiscellaneousCharge *Money `xml:"TaxesOrMiscellaneousCharge,omitempty"`

	// Specifies which kind of charge is being recorded in the preceding field.
	TaxesOrMiscellaneousChargeType *TaxesOrMiscellaneousChargeType `xml:"TaxesOrMiscellaneousChargeType,omitempty"`

	// Any packing costs that are associated with this shipment.
	PackingCosts *Money `xml:"PackingCosts,omitempty"`

	// Any handling costs that are associated with this shipment.
	HandlingCosts *Money `xml:"HandlingCosts,omitempty"`

	// Free-form text.
	SpecialInstructions string `xml:"SpecialInstructions,omitempty"`

	// Free-form text.
	DeclarationStatement string `xml:"DeclarationStatement,omitempty"`

	// Free-form text.
	PaymentTerms string `xml:"PaymentTerms,omitempty"`

	// The reason for the shipment. Note: SOLD is not a valid purpose for a Proforma Invoice.
	Purpose *PurposeOfShipmentType `xml:"Purpose,omitempty"`

	// Name of the International Expert that completed the Commercial Invoice different from Sender.
	OriginatorName string `xml:"OriginatorName,omitempty"`

	// Required for dutiable international Express or Ground shipments. This field is not applicable to an international PIB(document) or a non-document which does not require a Commercial Invoice.
	TermsOfSale string `xml:"TermsOfSale,omitempty"`
}
