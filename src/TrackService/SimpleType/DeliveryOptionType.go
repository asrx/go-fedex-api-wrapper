package SimpleType


// Specifies the different option types for delivery.
type DeliveryOptionType string

const (
	DeliveryOptionTypeINDIRECT_SIGNATURE_RELEASE DeliveryOptionType = "INDIRECT_SIGNATURE_RELEASE"

	DeliveryOptionTypeREDIRECT_TO_HOLD_AT_LOCATION DeliveryOptionType = "REDIRECT_TO_HOLD_AT_LOCATION"

	DeliveryOptionTypeREROUTE DeliveryOptionType = "REROUTE"

	DeliveryOptionTypeRESCHEDULE DeliveryOptionType = "RESCHEDULE"
)
