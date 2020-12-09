package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type FreightSpecialServicePayment struct {

	// Identifies the special service.
	SpecialService *SimpleType.ShipmentSpecialServiceType `xml:"SpecialService,omitempty"`

	// Indicates who will pay for the special service.
	PaymentType *SimpleType.FreightShipmentRoleType `xml:"PaymentType,omitempty"`
}
