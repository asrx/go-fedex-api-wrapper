package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type FreightBaseCharge struct {

	// Freight class for this line item.
	FreightClass *SimpleType.FreightClassType `xml:"FreightClass,omitempty"`

	// Effective freight class used for rating this line item.
	RatedAsClass *SimpleType.FreightClassType `xml:"RatedAsClass,omitempty"`

	// NMFC Code for commodity.
	NmfcCode string `xml:"NmfcCode,omitempty"`

	// Customer-provided description for this commodity or class line.
	Description string `xml:"Description,omitempty"`

	// Weight for this commodity or class line.
	Weight *Weight `xml:"Weight,omitempty"`

	// Rate or factor applied to this line item.
	ChargeRate *Money `xml:"ChargeRate,omitempty"`

	// Identifies the manner in which the chargeRate for this line item was applied.
	ChargeBasis *SimpleType.FreightChargeBasisType `xml:"ChargeBasis,omitempty"`

	// The net or extended charge for this line item.
	ExtendedAmount *Money `xml:"ExtendedAmount,omitempty"`
}
