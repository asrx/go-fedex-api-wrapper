package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackSelectionDetail struct {

	// The FedEx operating company (transportation) used for this package's delivery.
	CarrierCode *CarrierCodeType `xml:"CarrierCode,omitempty"`

	// Identifies operating transportation company that is the specific to the carrier code.
	OperatingCompany *OperatingCompanyType `xml:"OperatingCompany,omitempty"`

	// The type and value of the package identifier that is to be used to retrieve the tracking information for a package or group of packages.
	PackageIdentifier *TrackPackageIdentifier `xml:"PackageIdentifier,omitempty"`

	// Used to distinguish duplicate FedEx tracking numbers.
	TrackingNumberUniqueIdentifier string `xml:"TrackingNumberUniqueIdentifier,omitempty"`

	// To narrow the search to a period in time the ShipDateRangeBegin and ShipDateRangeEnd can be used to help eliminate duplicates.
	ShipDateRangeBegin string `xml:"ShipDateRangeBegin,omitempty"`

	// To narrow the search to a period in time the ShipDateRangeBegin and ShipDateRangeEnd can be used to help eliminate duplicates.
	ShipDateRangeEnd string `xml:"ShipDateRangeEnd,omitempty"`

	// For tracking by references information either the account number or destination postal code and country must be provided.
	ShipmentAccountNumber string `xml:"ShipmentAccountNumber,omitempty"`

	// Specifies the SPOD account number for the shipment being tracked.
	SecureSpodAccount string `xml:"SecureSpodAccount,omitempty"`

	// For tracking by references information either the account number or destination postal code and country must be provided.
	Destination *Address `xml:"Destination,omitempty"`

	// Specifies the details about how to retrieve the subsequent pages when there is more than one page in the TrackReply.
	PagingDetail *PagingDetail `xml:"PagingDetail,omitempty"`

	// The customer can specify a desired time out value for this particular tracking number.
	CustomerSpecifiedTimeOutValueInMilliseconds *uint `xml:"CustomerSpecifiedTimeOutValueInMilliseconds,omitempty"`
}
