package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type ServiceSubOptionDetail struct {

	// Identifies the type of Freight Guarantee applied, if FREIGHT_GUARANTEE is applied to the rate quote.
	FreightGuarantee *FreightGuaranteeType `xml:"FreightGuarantee,omitempty"`

	// Identifies the smartPostHubId used during rate quote, if SMART_POST_HUB_ID is a variable option on the rate request.
	SmartPostHubId string `xml:"SmartPostHubId,omitempty"`

	// Identifies the indicia used during rate quote, if SMART_POST_ALLOWED_INDICIA is a variable option on the rate request.
	SmartPostIndicia *SmartPostIndiciaType `xml:"SmartPostIndicia,omitempty"`
}
