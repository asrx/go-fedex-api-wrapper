package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ReturnEMailDetail struct {
	MerchantPhoneNumber string `xml:"MerchantPhoneNumber,omitempty"`

	// Identifies the allowed (merchant-authorized) special services which may be selected when the subsequent shipment is created. Only services represented in EMailLabelAllowedSpecialServiceType will be controlled by this list.
	AllowedSpecialServices []*SimpleType.ReturnEMailAllowedSpecialServiceType `xml:"AllowedSpecialServices,omitempty"`
}
