package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type EMailRecipient struct {

	// EMail address of the recipient.
	EmailAddress string `xml:"EmailAddress,omitempty"`

	// The relationship that the customer has to the pending shipment.
	Role *SimpleType.AccessorRoleType `xml:"Role,omitempty"`

	// Specifies how the email notification for the pending shipment need to be processed.
	OptionsRequested *EmailOptionsRequested `xml:"OptionsRequested,omitempty"`

	// Localization and language details specified by the recipient of the EMail.
	Localization *ComplexType.Localization `xml:"Localization,omitempty"`
}
