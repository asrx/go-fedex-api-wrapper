package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type PendingShipmentAccessorDetail struct {

	// Specifies the role of the user who is trying to access the pending shipment.
	Role *SimpleType.AccessorRoleType `xml:"Role,omitempty"`

	UserId string `xml:"UserId,omitempty"`

	Password string `xml:"Password,omitempty"`

	EmailLabelUrl string `xml:"EmailLabelUrl,omitempty"`
}