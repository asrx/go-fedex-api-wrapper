package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CompletedPackageDetail struct {
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	TrackingIds []*ComplexType.TrackingId `xml:"TrackingIds,omitempty"`

	// Used with request containing PACKAGE_GROUPS, to identify which group of identical packages was used to produce a reply item.
	GroupNumber *uint `xml:"GroupNumber,omitempty"`

	OversizeClass *SimpleType.OversizeClassType `xml:"OversizeClass,omitempty"`

	// All package-level rating data for this package, which may include data for multiple rate types.
	PackageRating *PackageRating `xml:"PackageRating,omitempty"`

	OperationalDetail *PackageOperationalDetail `xml:"OperationalDetail,omitempty"`

	Label *ShippingDocument `xml:"Label,omitempty"`

	// All package-level shipping documents (other than labels and barcodes). For use in loads after January, 2008.
	PackageDocuments []*ShippingDocument `xml:"PackageDocuments,omitempty"`

	// Specifies the information associated with this package that has COD special service in a ground shipment.
	CodReturnDetail *CodReturnPackageDetail `xml:"CodReturnDetail,omitempty"`

	// Actual signature option applied, to allow for cases in which the original value conflicted with other service features in the shipment.
	SignatureOption *SimpleType.SignatureOptionType `xml:"SignatureOption,omitempty"`

	DryIceWeight *ComplexType.Weight `xml:"DryIceWeight,omitempty"`

	// Documents the kinds and quantities of all hazardous commodities in the current package, using updated hazardous commodity description data.
	HazardousPackageDetail *CompletedHazardousPackageDetail `xml:"HazardousPackageDetail,omitempty"`
}
