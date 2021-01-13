package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"time"
)

type RequestedShipment struct {
	ShipTimestamp time.Time `xml:"ShipTimestamp,omitempty"`

	DropoffType *DropoffType `xml:"DropoffType,omitempty"`

	ServiceType *ServiceType `xml:"ServiceType,omitempty"`

	PackagingType *PackagingType `xml:"PackagingType,omitempty"`

	// The shipment variations for the current shipment expressed in key-value pairs.
	VariationOptions []*ShipmentVariationOptionDetail `xml:"VariationOptions,omitempty"`

	TotalWeight *Weight `xml:"TotalWeight,omitempty"`

	// Specifies the total declared value for carriage of the shipment. The declared value for carriage represents the maximum liability of FedEx in connection with a shipment, including, but not limited to, any loss, damage, delay, mis-delivery, nondelivery, misinformation, any failure to provide information, or mis-delivery of information relating to the shipment.
	TotalInsuredValue *Money `xml:"TotalInsuredValue,omitempty"`

	// This attribute indicates the currency the caller requests to have used in all returned monetary values (when a choice is possible).
	PreferredCurrency string `xml:"PreferredCurrency,omitempty"`

	// Specifies details about the entity responsible for the shipment.
	ShipmentAuthorizationDetail *ShipmentAuthorizationDetail `xml:"ShipmentAuthorizationDetail,omitempty"`

	Shipper *Party `xml:"Shipper,omitempty"`

	Recipient *Party `xml:"Recipient,omitempty"`

	RecipientLocationNumber string `xml:"RecipientLocationNumber,omitempty"`

	// Physical starting address for the shipment, if different from shipper's address.
	Origin *ContactAndAddress `xml:"Origin,omitempty"`

	// The sold-to party is used for customs clearance; for example, in support of US import customs rules. The need for this field could vary based on whether a sold-to party was specified on a consolidation.
	SoldTo *Party `xml:"SoldTo,omitempty"`

	ShippingChargesPayment *Payment `xml:"ShippingChargesPayment,omitempty"`

	SpecialServicesRequested *ShipmentSpecialServicesRequested `xml:"SpecialServicesRequested,omitempty"`

	ExpressFreightDetail *ExpressFreightDetail `xml:"ExpressFreightDetail,omitempty"`

	// Data applicable to shipments using FEDEX_FREIGHT_ECONOMY and FEDEX_FREIGHT_PRIORITY services.
	FreightShipmentDetail *FreightShipmentDetail `xml:"FreightShipmentDetail,omitempty"`

	// Used with Ground Home Delivery and Freight.
	DeliveryInstructions string `xml:"DeliveryInstructions,omitempty"`

	VariableHandlingChargeDetail *VariableHandlingChargeDetail `xml:"VariableHandlingChargeDetail,omitempty"`

	// Customs clearance data, used for both international and intra-country shipping.
	CustomsClearanceDetail *CustomsClearanceDetail `xml:"CustomsClearanceDetail,omitempty"`

	// For use in "process tag" transaction.
	PickupDetail *PickupDetail `xml:"PickupDetail,omitempty"`

	// Specifies the characteristics of a shipment pertaining to SmartPost services.
	SmartPostDetail *SmartPostShipmentDetail `xml:"SmartPostDetail,omitempty"`

	// 如果为真，则只有托运人/付款人可以看到这批货物。
	BlockInsightVisibility bool `xml:"BlockInsightVisibility,omitempty"`

	LabelSpecification *LabelSpecification `xml:"LabelSpecification,omitempty"`

	// Contains data used to create additional (non-label) shipping documents.
	ShippingDocumentSpecification *ShippingDocumentSpecification `xml:"ShippingDocumentSpecification,omitempty"`

	// 说明客户是否希望对这批货物报价以及报价的种类。回复也会受到发货和客户的其他数据的限制。
	RateRequestTypes []*RateRequestType `xml:"RateRequestTypes,omitempty"`

	// 指定客户是否希望在这批货物的价格报价中提供估计的关税和税收。只适用于国际运输。
	EdtRequestType *EdtRequestType `xml:"EdtRequestType,omitempty"`

	// 主单号
	// Only used with multiple-transaction shipments, to identify the master package in a multi-piece shipment.
	MasterTrackingId *TrackingId `xml:"MasterTrackingId,omitempty"`

	// The total number of packages in the entire shipment (even when the shipment spans multiple transactions.)
	PackageCount *uint `xml:"PackageCount,omitempty"`

	// Specifies which package-level data values are provided at the shipment-level only. The package-level data values types specified here will not be provided at the package-level.
	ShipmentOnlyFields []*ShipmentOnlyFieldsType `xml:"ShipmentOnlyFields,omitempty"`

	// 指定可在单次传送中重复使用多次的数据结构。
	ConfigurationData *ShipmentConfigurationData `xml:"ConfigurationData,omitempty"`

	// One or more package-attribute descriptions, each of which describes an individual package, a group of identical packages, or (for the total-piece-total-weight case) common characteristics all packages in the shipment.
	RequestedPackageLineItems []*RequestedPackageLineItem `xml:"RequestedPackageLineItems,omitempty"`
}
