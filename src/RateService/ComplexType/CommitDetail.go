package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/RateService/SimpleType"
)

type CommitDetail struct {

	// The Commodity applicable to this commitment.
	CommodityName string `xml:"CommodityName,omitempty"`

	// Deprecated: This field will be removed in a future DOM release. New code should use serviceDescription.serviceType instead.
	ServiceType *ServiceType `xml:"ServiceType,omitempty"`

	// Descriptions and alternate identifiers for a service.
	ServiceDescription *ServiceDescription `xml:"ServiceDescription,omitempty"`

	// Shows the specific combination of service options combined with the service type that produced this committment in the set returned to the caller.
	AppliedOptions []*ServiceOptionType `xml:"AppliedOptions,omitempty"`

	// Supporting detail for applied options identified in preceding field.
	AppliedSubOptions *ServiceSubOptionDetail `xml:"AppliedSubOptions,omitempty"`

	DerivedShipmentSignatureOption *SignatureOptionDetail `xml:"DerivedShipmentSignatureOption,omitempty"`

	DerivedPackageSignatureOptions []*SignatureOptionDetail `xml:"DerivedPackageSignatureOptions,omitempty"`

	DerivedOriginDetail *CleansedAddressAndLocationDetail `xml:"DerivedOriginDetail,omitempty"`

	DerivedDestinationDetail *CleansedAddressAndLocationDetail `xml:"DerivedDestinationDetail,omitempty"`

	// THe delivery commitment date/time. Express Only.
	CommitTimestamp string `xml:"CommitTimestamp,omitempty"`

	// The delivery commitment day of the week.
	DayOfWeek *DayOfWeekType `xml:"DayOfWeek,omitempty"`

	// The number of transit days; applies to Ground and LTL Freight; indicates minimum transit time for SmartPost.
	TransitTime *TransitTimeType `xml:"TransitTime,omitempty"`

	// Maximum number of transit days, for SmartPost shipments.
	MaximumTransitTime *TransitTimeType `xml:"MaximumTransitTime,omitempty"`

	// The service area code for the destination of this shipment. Express only.
	DestinationServiceArea string `xml:"DestinationServiceArea,omitempty"`

	// The address of the broker to be used for this shipment.
	BrokerAddress *Address `xml:"BrokerAddress,omitempty"`

	// The FedEx location identifier for the broker.
	BrokerLocationId string `xml:"BrokerLocationId,omitempty"`

	// The delivery commitment date/time the shipment will arrive at the border.
	BrokerCommitTimestamp string `xml:"BrokerCommitTimestamp,omitempty"`

	// The delivery commitment day of the week the shipment will arrive at the border.
	BrokerCommitDayOfWeek *DayOfWeekType `xml:"BrokerCommitDayOfWeek,omitempty"`

	// The number of days it will take for the shipment to make it from broker to destination
	BrokerToDestinationDays *uint `xml:"BrokerToDestinationDays,omitempty"`

	// The delivery commitment date for shipment served by GSP (Global Service Provider)
	ProofOfDeliveryDate string `xml:"ProofOfDeliveryDate,omitempty"`

	// The delivery commitment day of the week for the shipment served by GSP (Global Service Provider)
	ProofOfDeliveryDayOfWeek *DayOfWeekType `xml:"ProofOfDeliveryDayOfWeek,omitempty"`

	// Messages concerning the ability to provide an accurate delivery commitment on an International commit quote. These could be messages providing information about why a commitment could not be returned or a successful message such as "REQUEST COMPLETED"
	CommitMessages []*Notification `xml:"CommitMessages,omitempty"`

	// Messages concerning the delivery commitment on an International commit quote such as "0:00 A.M. IF NO CUSTOMS DELAY"
	DeliveryMessages []string `xml:"DeliveryMessages,omitempty"`

	// Information about why a shipment delivery is delayed and at what level (country/service etc.).
	DelayDetails []*DelayDetail `xml:"DelayDetails,omitempty"`

	DocumentContent *InternationalDocumentContentType `xml:"DocumentContent,omitempty"`

	// Required documentation for this shipment.
	RequiredDocuments []*RequiredShippingDocumentType `xml:"RequiredDocuments,omitempty"`

	// Freight origin and destination city center information and total distance between origin and destination city centers.
	FreightCommitDetail *FreightCommitDetail `xml:"FreightCommitDetail,omitempty"`
}
