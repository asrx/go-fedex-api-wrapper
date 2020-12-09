package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CodDetail struct {
	CodCollectionAmount *Money `xml:"CodCollectionAmount,omitempty"`

	// Specifies the details of the charges are to be added to the COD collect amount.
	AddTransportationChargesDetail *CodAddTransportationChargesDetail `xml:"AddTransportationChargesDetail,omitempty"`

	// Identifies the type of funds FedEx should collect upon package delivery
	CollectionType *SimpleType.CodCollectionType `xml:"CollectionType,omitempty"`

	// For Express this is the descriptive data that is used for the recipient of the FedEx Letter containing the COD payment. For Ground this is the descriptive data for the party to receive the payment that prints the COD receipt.
	CodRecipient *Party `xml:"CodRecipient,omitempty"`

	// When the FedEx COD payment type is not CASH, indicates the contact and address of the financial institution used to service the payment of the COD.
	FinancialInstitutionContactAndAddress *ContactAndAddress `xml:"FinancialInstitutionContactAndAddress,omitempty"`

	// Specifies the name of person or company receiving the secured/unsecured funds payment
	RemitToName string `xml:"RemitToName,omitempty"`

	// Indicates which type of reference information to include on the COD return shipping label.
	ReferenceIndicator *SimpleType.CodReturnReferenceIndicatorType `xml:"ReferenceIndicator,omitempty"`

	// Only used with multi-piece COD shipments sent in multiple transactions. Required on last transaction only.
	ReturnTrackingId *TrackingId `xml:"ReturnTrackingId,omitempty"`
}
