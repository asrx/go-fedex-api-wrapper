package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShipmentManifestDetail struct {

	// This identifies which customer reference field contains the manifest ID.
	ManifestReferenceType *SimpleType.CustomerReferenceType `xml:"ManifestReferenceType,omitempty"`
}
