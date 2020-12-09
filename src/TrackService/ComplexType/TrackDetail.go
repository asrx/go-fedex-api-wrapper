package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackDetail struct {

	// To report soft error on an individual track detail.
	Notification *Notification `xml:"Notification,omitempty"`

	// The FedEx package identifier.
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	Barcode *StringBarcode `xml:"Barcode,omitempty"`

	// When duplicate tracking numbers exist this data is returned with summary information for each of the duplicates. The summary information is used to determine which of the duplicates the intended tracking number is. This identifier is used on a subsequent track request to retrieve the tracking data for the desired tracking number.
	TrackingNumberUniqueIdentifier string `xml:"TrackingNumberUniqueIdentifier,omitempty"`

	// Specifies details about the status of the shipment being tracked.
	StatusDetail *TrackStatusDetail `xml:"StatusDetail,omitempty"`

	// Notifications to the end user that provide additional information relevant to the tracked shipment. For example, a notification may indicate that a change in behavior has occurred.
	InformationNotes []*TrackInformationNoteDetail `xml:"InformationNotes,omitempty"`

	CustomerExceptionRequests []*CustomerExceptionRequestDetail `xml:"CustomerExceptionRequests,omitempty"`

	// Used to report the status of a piece of a multiple piece shipment which is no longer traveling with the rest of the packages in the shipment or has not been accounted for.
	Reconciliation *TrackReconciliation `xml:"Reconciliation,omitempty"`

	// Used to convey information such as. 1. FedEx has received information about a package but has not yet taken possession of it. 2. FedEx has handed the package off to a third party for final delivery. 3. The package delivery has been cancelled
	ServiceCommitMessage string `xml:"ServiceCommitMessage,omitempty"`

	DestinationServiceArea string `xml:"DestinationServiceArea,omitempty"`

	DestinationServiceAreaDescription string `xml:"DestinationServiceAreaDescription,omitempty"`

	// Identifies a FedEx operating company (transportation).
	CarrierCode *CarrierCodeType `xml:"CarrierCode,omitempty"`

	// Identifies operating transportation company that is the specific to the carrier code.
	OperatingCompany *OperatingCompanyType `xml:"OperatingCompany,omitempty"`

	// Specifies a detailed description about the carrier or the operating company.
	OperatingCompanyOrCarrierDescription string `xml:"OperatingCompanyOrCarrierDescription,omitempty"`

	// If the package was interlined to a cartage agent, this is the name of the cartage agent. (Returned for CSR SL only.)
	CartageAgentCompanyName string `xml:"CartageAgentCompanyName,omitempty"`

	// Specifies the FXO production centre contact and address.
	ProductionLocationContactAndAddress *ContactAndAddress `xml:"ProductionLocationContactAndAddress,omitempty"`

	OtherIdentifiers []*TrackOtherIdentifierDetail `xml:"OtherIdentifiers,omitempty"`

	// (Returned for CSR SL only.)
	FormId string `xml:"FormId,omitempty"`

	// Specifies details about service such as service description and type.
	Service *TrackServiceDescriptionDetail `xml:"Service,omitempty"`

	// The weight of this package.
	PackageWeight *Weight `xml:"PackageWeight,omitempty"`

	// Physical dimensions of the package.
	PackageDimensions *Dimensions `xml:"PackageDimensions,omitempty"`

	// The dimensional weight of the package.
	PackageDimensionalWeight *Weight `xml:"PackageDimensionalWeight,omitempty"`

	// The weight of the entire shipment.
	ShipmentWeight *Weight `xml:"ShipmentWeight,omitempty"`

	// Retained for legacy compatibility only.
	Packaging string `xml:"Packaging,omitempty"`

	// Strict representation of the Packaging type (e.g. FEDEX_BOX, YOUR_PACKAGING).
	PackagingType *PackagingType `xml:"PackagingType,omitempty"`

	PhysicalPackagingType *PhysicalPackagingType `xml:"PhysicalPackagingType,omitempty"`

	// The sequence number of this package in a shipment. This would be 2 if it was package number 2 of 4.
	PackageSequenceNumber *uint `xml:"PackageSequenceNumber,omitempty"`

	// The number of packages in this shipment.
	PackageCount *uint `xml:"PackageCount,omitempty"`

	// FOR FEDEX INTERNAL USE ONLY: Specifies the software id of the device that was used to create this tracked shipment.
	CreatorSoftwareId string `xml:"CreatorSoftwareId,omitempty"`

	// Specifies the details about the SPOC details.
	Charges []*TrackChargeDetail `xml:"Charges,omitempty"`

	NickName string `xml:"NickName,omitempty"`

	Notes string `xml:"Notes,omitempty"`

	Attributes []*TrackDetailAttributeType `xml:"Attributes,omitempty"`

	ShipmentContents []*ContentRecord `xml:"ShipmentContents,omitempty"`

	PackageContents []string `xml:"PackageContents,omitempty"`

	ClearanceLocationCode string `xml:"ClearanceLocationCode,omitempty"`

	Commodities []*Commodity `xml:"Commodities,omitempty"`

	ReturnDetail *TrackReturnDetail `xml:"ReturnDetail,omitempty"`

	// Specifies the reason for return.
	CustomsOptionDetails []*CustomsOptionDetail `xml:"CustomsOptionDetails,omitempty"`

	AdvanceNotificationDetail *TrackAdvanceNotificationDetail `xml:"AdvanceNotificationDetail,omitempty"`

	// List of special handlings that applied to this package.
	SpecialHandlings []*TrackSpecialHandling `xml:"SpecialHandlings,omitempty"`

	// Specifies the details about the payments for the shipment being tracked.
	Payments []*TrackPayment `xml:"Payments,omitempty"`

	// (Returned for CSR SL only.)
	Shipper *Contact `xml:"Shipper,omitempty"`

	// Indicates last-known possession of package (Returned for CSR SL only.)
	PossessionStatus *TrackPossessionStatusType `xml:"PossessionStatus,omitempty"`

	ShipperAddress *Address `xml:"ShipperAddress,omitempty"`

	// The address of the FedEx pickup location/facility.
	OriginLocationAddress *Address `xml:"OriginLocationAddress,omitempty"`

	// (Returned for CSR SL only.)
	OriginStationId string `xml:"OriginStationId,omitempty"`

	DatesOrTimes []*TrackingDateOrTimestamp `xml:"DatesOrTimes,omitempty"`

	// The distance from the origin to the destination. Returned for Custom Critical shipments.
	TotalTransitDistance *Distance `xml:"TotalTransitDistance,omitempty"`

	// Total distance package still has to travel. Returned for Custom Critical shipments.
	DistanceToDestination *Distance `xml:"DistanceToDestination,omitempty"`

	// Provides additional details about package delivery.
	SpecialInstructions []*TrackSpecialInstruction `xml:"SpecialInstructions,omitempty"`

	// (Returned for CSR SL only.)
	Recipient *Contact `xml:"Recipient,omitempty"`

	// This is the latest updated destination address.
	LastUpdatedDestinationAddress *Address `xml:"LastUpdatedDestinationAddress,omitempty"`

	// The address this package is to be (or has been) delivered.
	DestinationAddress *Address `xml:"DestinationAddress,omitempty"`

	HoldAtLocationContact *Contact `xml:"HoldAtLocationContact,omitempty"`

	// The address this package is requested to placed on hold.
	HoldAtLocationAddress *Address `xml:"HoldAtLocationAddress,omitempty"`

	// (Returned for CSR SL only.)
	DestinationStationId string `xml:"DestinationStationId,omitempty"`

	// The address of the FedEx delivery location/facility.
	DestinationLocationAddress *Address `xml:"DestinationLocationAddress,omitempty"`

	DestinationLocationType *FedExLocationType `xml:"DestinationLocationType,omitempty"`

	DestinationLocationTimeZoneOffset string `xml:"DestinationLocationTimeZoneOffset,omitempty"`

	// Actual address where package was delivered. Differs from destinationAddress, which indicates where the package was to be delivered; This field tells where delivery actually occurred (next door, at station, etc.)
	ActualDeliveryAddress *Address `xml:"ActualDeliveryAddress,omitempty"`

	// Identifies the method of office order delivery.
	OfficeOrderDeliveryMethod *OfficeOrderDeliveryMethodType `xml:"OfficeOrderDeliveryMethod,omitempty"`

	// Strict text indicating the delivery location at the delivered to address.
	DeliveryLocationType *TrackDeliveryLocationType `xml:"DeliveryLocationType,omitempty"`

	// User/screen friendly representation of the DeliveryLocationType (delivery location at the delivered to address). Can be returned in localized text.
	DeliveryLocationDescription string `xml:"DeliveryLocationDescription,omitempty"`

	// Specifies the number of delivery attempts made to deliver the shipment.
	DeliveryAttempts *uint `xml:"DeliveryAttempts,omitempty"`

	// This is either the name of the person that signed for the package or "Signature not requested" or "Signature on file".
	DeliverySignatureName string `xml:"DeliverySignatureName,omitempty"`

	// Specifies the details about the count of the packages delivered at the delivery location and the count of the packages at the origin.
	PieceCountVerificationDetails []*PieceCountVerificationDetail `xml:"PieceCountVerificationDetails,omitempty"`

	// Specifies the total number of unique addresses on the CRNs in a consolidation.
	TotalUniqueAddressCountInConsolidation *uint `xml:"TotalUniqueAddressCountInConsolidation,omitempty"`

	AvailableImages []*AvailableImagesDetail `xml:"AvailableImages,omitempty"`

	Signature *SignatureImageDetail `xml:"Signature,omitempty"`

	NotificationEventsAvailable []*NotificationEventType `xml:"NotificationEventsAvailable,omitempty"`

	// Returned for cargo shipments only when they are currently split across vehicles.
	SplitShipmentParts []*TrackSplitShipmentPart `xml:"SplitShipmentParts,omitempty"`

	// Specifies the details about the eligibility for different delivery options.
	DeliveryOptionEligibilityDetails []*DeliveryOptionEligibilityDetail `xml:"DeliveryOptionEligibilityDetails,omitempty"`

	// Event information for a tracking number.
	Events []*TrackEvent `xml:"Events,omitempty"`
}
