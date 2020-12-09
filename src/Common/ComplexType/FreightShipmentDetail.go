package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type FreightShipmentDetail struct {

	// Account number used with FEDEX_FREIGHT service.
	FedExFreightAccountNumber string `xml:"FedExFreightAccountNumber,omitempty"`

	// Used for validating FedEx Freight account number and (optionally) identifying third party payment on the bill of lading.
	FedExFreightBillingContactAndAddress *ContactAndAddress `xml:"FedExFreightBillingContactAndAddress,omitempty"`

	// Used in connection with "Send Bill To" (SBT) identification of customer's account used for billing.
	AlternateBilling *Party `xml:"AlternateBilling,omitempty"`

	// Indicates the role of the party submitting the transaction.
	Role *SimpleType.FreightShipmentRoleType `xml:"Role,omitempty"`

	// Designates the terms of the "collect" payment for a Freight Shipment.
	CollectTermsType *SimpleType.FreightCollectTermsType `xml:"CollectTermsType,omitempty"`

	// Identifies the declared value for the shipment
	DeclaredValuePerUnit *Money `xml:"DeclaredValuePerUnit,omitempty"`

	// Identifies the declared value units corresponding to the above defined declared value
	DeclaredValueUnits string `xml:"DeclaredValueUnits,omitempty"`

	LiabilityCoverageDetail *LiabilityCoverageDetail `xml:"LiabilityCoverageDetail,omitempty"`

	// Identifiers for promotional discounts offered to customers.
	Coupons []string `xml:"Coupons,omitempty"`

	// Total number of individual handling units in the entire shipment (for unit pricing).
	TotalHandlingUnits *uint `xml:"TotalHandlingUnits,omitempty"`

	// Estimated discount rate provided by client for unsecured rate quote.
	ClientDiscountPercent float64 `xml:"ClientDiscountPercent,omitempty"`

	// Total weight of pallets used in shipment.
	PalletWeight *Weight `xml:"PalletWeight,omitempty"`

	// Overall shipment dimensions.
	ShipmentDimensions *Dimensions `xml:"ShipmentDimensions,omitempty"`

	// Description for the shipment.
	Comment string `xml:"Comment,omitempty"`

	// Specifies which party will pay surcharges for any special services which support split billing.
	SpecialServicePayments []*FreightSpecialServicePayment `xml:"SpecialServicePayments,omitempty"`

	HazardousMaterialsOfferor string `xml:"HazardousMaterialsOfferor,omitempty"`

	// Details of the commodities in the shipment.
	LineItems []*FreightShipmentLineItem `xml:"LineItems,omitempty"`
}
