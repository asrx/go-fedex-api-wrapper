package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ShipmentSpecialServicesRequested struct {

	// The types of all special services requested for the enclosing shipment (or other shipment-level transaction).
	SpecialServiceTypes []*ShipmentSpecialServiceType `xml:"SpecialServiceTypes,omitempty"`

	CodDetail *CodDetail `xml:"CodDetail,omitempty"`

	DeliveryOnInvoiceAcceptanceDetail *DeliveryOnInvoiceAcceptanceDetail `xml:"DeliveryOnInvoiceAcceptanceDetail,omitempty"`

	HoldAtLocationDetail *HoldAtLocationDetail `xml:"HoldAtLocationDetail,omitempty"`

	// This replaces eMailNotificationDetail
	EventNotificationDetail *ShipmentEventNotificationDetail `xml:"EventNotificationDetail,omitempty"`

	ReturnShipmentDetail *ReturnShipmentDetail `xml:"ReturnShipmentDetail,omitempty"`

	// This field should be populated for pending shipments (e.g. e-mail label) It is required by a PENDING_SHIPMENT special service type.
	PendingShipmentDetail *PendingShipmentDetail `xml:"PendingShipmentDetail,omitempty"`

	InternationalControlledExportDetail *InternationalControlledExportDetail `xml:"InternationalControlledExportDetail,omitempty"`

	InternationalTrafficInArmsRegulationsDetail *InternationalTrafficInArmsRegulationsDetail `xml:"InternationalTrafficInArmsRegulationsDetail,omitempty"`

	ShipmentDryIceDetail *ShipmentDryIceDetail `xml:"ShipmentDryIceDetail,omitempty"`

	HomeDeliveryPremiumDetail *HomeDeliveryPremiumDetail `xml:"HomeDeliveryPremiumDetail,omitempty"`

	FlatbedTrailerDetail *FlatbedTrailerDetail `xml:"FlatbedTrailerDetail,omitempty"`

	FreightGuaranteeDetail *FreightGuaranteeDetail `xml:"FreightGuaranteeDetail,omitempty"`

	// Electronic Trade document references.
	EtdDetail *EtdDetail `xml:"EtdDetail,omitempty"`

	// Specification for date or range of dates on which delivery is to be attempted.
	CustomDeliveryWindowDetail *CustomDeliveryWindowDetail `xml:"CustomDeliveryWindowDetail,omitempty"`
}
