package SimpleType


// These values control the optional features of service that may be combined in a commitment/rate comparision transaction.
type ServiceOptionType string

const (
	ServiceOptionTypeFEDEX_ONE_RATE ServiceOptionType = "FEDEX_ONE_RATE"

	ServiceOptionTypeFREIGHT_GUARANTEE ServiceOptionType = "FREIGHT_GUARANTEE"

	ServiceOptionTypeSATURDAY_DELIVERY ServiceOptionType = "SATURDAY_DELIVERY"

	ServiceOptionTypeSMART_POST_ALLOWED_INDICIA ServiceOptionType = "SMART_POST_ALLOWED_INDICIA"

	ServiceOptionTypeSMART_POST_HUB_ID ServiceOptionType = "SMART_POST_HUB_ID"
)
