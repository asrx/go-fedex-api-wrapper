package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CompletedShipmentDetail struct {
	UsDomestic bool `xml:"UsDomestic,omitempty"`

	CarrierCode *CarrierCodeType `xml:"CarrierCode,omitempty"`

	MasterTrackingId *TrackingId `xml:"MasterTrackingId,omitempty"`

	// DEPRECATED as of 201801: Use serviceDescription instead.
	ServiceTypeDescription string `xml:"ServiceTypeDescription,omitempty"`

	ServiceDescription *ServiceDescription `xml:"ServiceDescription,omitempty"`

	PackagingDescription string `xml:"PackagingDescription,omitempty"`

	OperationalDetail *ShipmentOperationalDetail `xml:"OperationalDetail,omitempty"`

	// Only used with pending shipments.
	AccessDetail *PendingShipmentAccessDetail `xml:"AccessDetail,omitempty"`

	// Only used in the reply to tag requests.
	TagDetail *CompletedTagDetail `xml:"TagDetail,omitempty"`

	SmartPostDetail *CompletedSmartPostDetail `xml:"SmartPostDetail,omitempty"`

	// Computed shipment level information about hazarous commodities.
	HazardousShipmentDetail *CompletedHazardousShipmentDetail `xml:"HazardousShipmentDetail,omitempty"`

	// All shipment-level rating data for this shipment, which may include data for multiple rate types.
	ShipmentRating *ShipmentRating `xml:"ShipmentRating,omitempty"`

	// Returns the default holding location information when HOLD_AT_LOCATION special service is requested and the client does not specify the hold location address.
	CompletedHoldAtLocationDetail *CompletedHoldAtLocationDetail `xml:"CompletedHoldAtLocationDetail,omitempty"`

	// Returns any defaults or updates applied to RequestedShipment.exportDetail.exportComplianceStatement.
	ExportComplianceStatement string `xml:"ExportComplianceStatement,omitempty"`

	// This specifies what rules or requirements for documents are applicable for this shipment. This may identify required or prohibited documents.
	DocumentRequirements *DocumentRequirementsDetail `xml:"DocumentRequirements,omitempty"`

	CompletedEtdDetail *CompletedEtdDetail `xml:"CompletedEtdDetail,omitempty"`

	// All shipment-level shipping documents (other than labels and barcodes).
	ShipmentDocuments []*ShippingDocument `xml:"ShipmentDocuments,omitempty"`

	AssociatedShipments []*AssociatedShipmentDetail `xml:"AssociatedShipments,omitempty"`

	CompletedCodDetail *CompletedCodDetail `xml:"CompletedCodDetail,omitempty"`

	CompletedPackageDetails []*CompletedPackageDetail `xml:"CompletedPackageDetails,omitempty"`
}
