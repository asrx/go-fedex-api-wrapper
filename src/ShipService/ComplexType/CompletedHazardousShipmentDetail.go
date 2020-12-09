package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type CompletedHazardousShipmentDetail struct {
	HazardousSummaryDetail *CompletedHazardousSummaryDetail `xml:"HazardousSummaryDetail,omitempty"`

	DryIceDetail *ComplexType.ShipmentDryIceDetail `xml:"DryIceDetail,omitempty"`

	// This contains the ADR License information, which identifies the license number and ADR category under which the customer is allowed to ship.
	AdrLicense *AdrLicenseDetail `xml:"AdrLicense,omitempty"`
}
