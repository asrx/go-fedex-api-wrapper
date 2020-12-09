package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackPackageIdentifier struct {

	// The type of the Value to be used to retrieve tracking information for a package (e.g. SHIPPER_REFERENCE, PURCHASE_ORDER, TRACKING_NUMBER_OR_DOORTAG, etc..) .
	Type *SimpleType.TrackIdentifierType `xml:"Type,omitempty"`

	// The value to be used to retrieve tracking information for a package.
	Value string `xml:"Value,omitempty"`
}
