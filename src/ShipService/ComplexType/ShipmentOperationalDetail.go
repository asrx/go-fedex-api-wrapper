package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type ShipmentOperationalDetail struct {
	UrsaPrefixCode string `xml:"UrsaPrefixCode,omitempty"`

	UrsaSuffixCode string `xml:"UrsaSuffixCode,omitempty"`

	OriginLocationId string `xml:"OriginLocationId,omitempty"`

	OriginLocationNumber int32 `xml:"OriginLocationNumber,omitempty"`

	OriginServiceArea string `xml:"OriginServiceArea,omitempty"`

	DestinationLocationId string `xml:"DestinationLocationId,omitempty"`

	DestinationLocationNumber int32 `xml:"DestinationLocationNumber,omitempty"`

	DestinationServiceArea string `xml:"DestinationServiceArea,omitempty"`

	// This is the state of the destination location ID, and is not necessarily the same as the postal state.
	DestinationLocationStateOrProvinceCode string `xml:"DestinationLocationStateOrProvinceCode,omitempty"`

	// Expected/estimated date of delivery.
	DeliveryDate string `xml:"DeliveryDate,omitempty"`

	// Expected/estimated day of week of delivery.
	DeliveryDay *DayOfWeekType `xml:"DeliveryDay,omitempty"`

	// Delivery time, as published in Service Guide.
	PublishedDeliveryTime string `xml:"PublishedDeliveryTime,omitempty"`

	// Committed date of delivery.
	CommitDate string `xml:"CommitDate,omitempty"`

	// Committed day of week of delivery.
	CommitDay *DayOfWeekType `xml:"CommitDay,omitempty"`

	// Standard transit time per origin, destination, and service.
	TransitTime *TransitTimeType `xml:"TransitTime,omitempty"`

	// Maximum expected transit time
	MaximumTransitTime *TransitTimeType `xml:"MaximumTransitTime,omitempty"`

	// Transit time based on customer eligibility.
	CustomTransitTime *TransitTimeType `xml:"CustomTransitTime,omitempty"`

	// Indicates that this shipment is not eligible for money back guarantee.
	IneligibleForMoneyBackGuarantee bool `xml:"IneligibleForMoneyBackGuarantee,omitempty"`

	// FedEx Ground delivery features for which this shipment may be eligible.
	DeliveryEligibilities []*SimpleType.GroundDeliveryEligibilityType `xml:"DeliveryEligibilities,omitempty"`

	// Text describing planned delivery.
	AstraPlannedServiceLevel string `xml:"AstraPlannedServiceLevel,omitempty"`

	AstraDescription string `xml:"AstraDescription,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	StateOrProvinceCode string `xml:"StateOrProvinceCode,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	AirportId string `xml:"AirportId,omitempty"`

	ServiceCode string `xml:"ServiceCode,omitempty"`

	PackagingCode string `xml:"PackagingCode,omitempty"`

	Scac string `xml:"Scac,omitempty"`
}
