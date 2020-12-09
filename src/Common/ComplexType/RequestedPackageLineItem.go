package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RequestedPackageLineItem struct {

	// Used only with INDIVIDUAL_PACKAGE, as a unique identifier of each requested package.
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	// Used only with PACKAGE_GROUPS, as a unique identifier of each group of identical packages.
	GroupNumber *uint `xml:"GroupNumber,omitempty"`

	// Used only with PACKAGE_GROUPS, as a count of packages within a group of identical packages.
	GroupPackageCount *uint `xml:"GroupPackageCount,omitempty"`

	VariableHandlingChargeDetail *VariableHandlingChargeDetail `xml:"VariableHandlingChargeDetail,omitempty"`

	// Specifies the declared value for carriage of the package. The declared value for carriage represents the maximum liability of FedEx in connection with a shipment, including, but not limited to, any loss, damage, delay, mis-delivery, nondelivery, misinformation, any failure to provide information, or mis-delivery of information relating to the package. This field is only used for INDIVIDUAL_PACKAGES and PACKAGE_GROUPS. Ignored for PACKAGE_SUMMARY, in which case totalInsuredValue and packageCount on the shipment will be used to determine this value.
	InsuredValue *Money `xml:"InsuredValue,omitempty"`

	// Only used for INDIVIDUAL_PACKAGES and PACKAGE_GROUPS. Ignored for PACKAGE_SUMMARY, in which case total weight and packageCount on the shipment will be used to determine this value.
	Weight *Weight `xml:"Weight,omitempty"`

	Dimensions *Dimensions `xml:"Dimensions,omitempty"`

	// Provides additional detail on how the customer has physically packaged this item. As of June 2009, required for packages moving under international and SmartPost services.
	PhysicalPackaging *PhysicalPackagingType `xml:"PhysicalPackaging,omitempty"`

	AssociatedFreightLineItems []*AssociatedFreightLineItemDetail `xml:"AssociatedFreightLineItems,omitempty"`

	// Human-readable text describing the package.
	ItemDescription string `xml:"ItemDescription,omitempty"`

	// Human-readable text describing the contents of the package to be used for clearance purposes.
	ItemDescriptionForClearance string `xml:"ItemDescriptionForClearance,omitempty"`

	CustomerReferences *CustomerReference `xml:"CustomerReferences,omitempty"`

	SpecialServicesRequested *PackageSpecialServicesRequested `xml:"SpecialServicesRequested,omitempty"`

	// Only used for INDIVIDUAL_PACKAGES and PACKAGE_GROUPS.
	ContentRecords []*ContentRecord `xml:"ContentRecords,omitempty"`
}
