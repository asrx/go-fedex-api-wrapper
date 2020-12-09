package SimpleType

type TrackDeliveryOptionType string

const (
	TrackDeliveryOptionTypeAPPOINTMENT TrackDeliveryOptionType = "APPOINTMENT"

	TrackDeliveryOptionTypeDATE_CERTAIN TrackDeliveryOptionType = "DATE_CERTAIN"

	TrackDeliveryOptionTypeELECTRONIC_SIGNATURE_RELEASE TrackDeliveryOptionType = "ELECTRONIC_SIGNATURE_RELEASE"

	TrackDeliveryOptionTypeEVENING TrackDeliveryOptionType = "EVENING"

	TrackDeliveryOptionTypeREDIRECT_TO_HOLD_AT_LOCATION TrackDeliveryOptionType = "REDIRECT_TO_HOLD_AT_LOCATION"

	TrackDeliveryOptionTypeREROUTE TrackDeliveryOptionType = "REROUTE"
)
