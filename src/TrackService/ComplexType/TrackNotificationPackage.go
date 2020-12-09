package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type TrackNotificationPackage struct {

	// FedEx assigned identifier for a package/shipment.
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	// When duplicate tracking numbers exist this data is returned with summary information for each of the duplicates. The summary information is used to determine which of the duplicates the intended tracking number is. This identifier is used on a subsequent track request to retrieve the tracking data for the desired tracking number.
	TrackingNumberUniqueIdentifiers string `xml:"TrackingNumberUniqueIdentifiers,omitempty"`

	// Identification of a FedEx operating company (transportation).
	CarrierCode *SimpleType.CarrierCodeType `xml:"CarrierCode,omitempty"`

	// The date the package was shipped (tendered to FedEx).
	ShipDate string `xml:"ShipDate,omitempty"`

	// The destination address of this package. Only city, state/province, and country are returned.
	Destination *ComplexType.Address `xml:"Destination,omitempty"`

	// Options available for a tracking notification recipient.
	RecipientDetails []*TrackNotificationRecipientDetail `xml:"RecipientDetails,omitempty"`
}
