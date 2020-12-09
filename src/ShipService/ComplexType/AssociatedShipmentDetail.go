package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type AssociatedShipmentDetail struct {
	Type *AssociatedShipmentType `xml:"Type,omitempty"`

	Sender *Party `xml:"Sender,omitempty"`

	Recipient *Party `xml:"Recipient,omitempty"`

	ServiceType *ServiceType `xml:"ServiceType,omitempty"`

	PackagingType *PackagingType `xml:"PackagingType,omitempty"`

	// Specifies the tracking id for the payment on the return.
	TrackingId *TrackingId `xml:"TrackingId,omitempty"`

	// Specifies additional customer reference data about the associated shipment.
	CustomerReferences []*CustomerReference `xml:"CustomerReferences,omitempty"`

	// Specifies shipment level operational information.
	ShipmentOperationalDetail *ShipmentOperationalDetail `xml:"ShipmentOperationalDetail,omitempty"`

	// Specifies package level operational information on the associated shipment. This information is not tied to an individual outbound package.
	PackageOperationalDetail *PackageOperationalDetail `xml:"PackageOperationalDetail,omitempty"`

	Label *ShippingDocument `xml:"Label,omitempty"`
}

