package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RateReplyDetail struct {

	// Deprecated: This field will be removed in a future DOM release. New code should use serviceDescription.serviceType instead.
	ServiceType *ServiceType `xml:"ServiceType,omitempty"`

	// Descriptions and alternate identifiers for a service.
	ServiceDescription *ServiceDescription `xml:"ServiceDescription,omitempty"`

	PackagingType *PackagingType `xml:"PackagingType,omitempty"`

	// Shows the specific combination of service options combined with the service type that produced this committment in the set returned to the caller.
	AppliedOptions []*ServiceOptionType `xml:"AppliedOptions,omitempty"`

	// Supporting detail for applied options identified in preceding field.
	AppliedSubOptions *ServiceSubOptionDetail `xml:"AppliedSubOptions,omitempty"`

	DeliveryStation string `xml:"DeliveryStation,omitempty"`

	DeliveryDayOfWeek *DayOfWeekType `xml:"DeliveryDayOfWeek,omitempty"`

	DeliveryTimestamp string `xml:"DeliveryTimestamp,omitempty"`

	CommitDetails []*CommitDetail `xml:"CommitDetails,omitempty"`

	DestinationAirportId string `xml:"DestinationAirportId,omitempty"`

	IneligibleForMoneyBackGuarantee bool `xml:"IneligibleForMoneyBackGuarantee,omitempty"`

	// Not populated by FAST service in Jan07.
	OriginServiceArea string `xml:"OriginServiceArea,omitempty"`

	// Not populated by FAST service in Jan07.
	DestinationServiceArea string `xml:"DestinationServiceArea,omitempty"`

	// Not populated by FAST service in Jan07.
	TransitTime *TransitTimeType `xml:"TransitTime,omitempty"`

	// Maximum expected transit time
	MaximumTransitTime *TransitTimeType `xml:"MaximumTransitTime,omitempty"`

	// Not populated by FAST service in Jan07. Actual signature option applied, to allow for cases in wihch the original value conflicted with other service features in the shipment.
	SignatureOption *SignatureOptionType `xml:"SignatureOption,omitempty"`

	ActualRateType *ReturnedRateType `xml:"ActualRateType,omitempty"`

	// Each element contains all rate data for a single rate type.
	RatedShipmentDetails []*RatedShipmentDetail `xml:"RatedShipmentDetails,omitempty"`
}
