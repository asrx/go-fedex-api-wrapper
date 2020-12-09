package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CustomsClearanceDetail struct {
	Brokers []*BrokerDetail `xml:"Brokers,omitempty"`

	// Interacts both with properties of the shipment and contractual relationship with the shipper.
	ClearanceBrokerage *SimpleType.ClearanceBrokerageType `xml:"ClearanceBrokerage,omitempty"`

	CustomsOptions *CustomsOptionDetail `xml:"CustomsOptions,omitempty"`

	ImporterOfRecord *Party `xml:"ImporterOfRecord,omitempty"`

	// Specifies how the recipient is identified for customs purposes; the requirements on this information vary with destination country.
	RecipientCustomsId *RecipientCustomsId `xml:"RecipientCustomsId,omitempty"`

	DutiesPayment *Payment `xml:"DutiesPayment,omitempty"`

	DocumentContent *SimpleType.InternationalDocumentContentType `xml:"DocumentContent,omitempty"`

	CustomsValue *Money `xml:"CustomsValue,omitempty"`

	// Identifies responsibilities with respect to loss, damage, etc.
	FreightOnValue *SimpleType.FreightOnValueType `xml:"FreightOnValue,omitempty"`

	// Documents amount paid to third party for coverage of shipment content.
	InsuranceCharges *Money `xml:"InsuranceCharges,omitempty"`

	PartiesToTransactionAreRelated bool `xml:"PartiesToTransactionAreRelated,omitempty"`

	CommercialInvoice *CommercialInvoice `xml:"CommercialInvoice,omitempty"`

	Commodities []*Commodity `xml:"Commodities,omitempty"`

	ExportDetail *ExportDetail `xml:"ExportDetail,omitempty"`

	RegulatoryControls []*SimpleType.RegulatoryControlType `xml:"RegulatoryControls,omitempty"`
}
