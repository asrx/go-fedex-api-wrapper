package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type HoldAtLocationDetail struct {

	// Contact phone number for recipient of shipment.
	PhoneNumber string `xml:"PhoneNumber,omitempty"`

	// Contact and address of FedEx facility at which shipment is to be held.
	LocationContactAndAddress *ContactAndAddress `xml:"LocationContactAndAddress,omitempty"`

	// Type of facility at which package/shipment is to be held.
	LocationType *SimpleType.FedExLocationType `xml:"LocationType,omitempty"`

	// Location identification (for facilities identified by an alphanumeric location code).
	LocationId string `xml:"LocationId,omitempty"`

	// Location identification (for facilities identified by an numeric location code).
	LocationNumber int32 `xml:"LocationNumber,omitempty"`
}
