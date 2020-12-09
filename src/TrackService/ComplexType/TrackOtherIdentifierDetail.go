package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type TrackOtherIdentifierDetail struct {
	PackageIdentifier *TrackPackageIdentifier `xml:"PackageIdentifier,omitempty"`

	TrackingNumberUniqueIdentifier string `xml:"TrackingNumberUniqueIdentifier,omitempty"`

	CarrierCode *SimpleType.CarrierCodeType `xml:"CarrierCode,omitempty"`
}
