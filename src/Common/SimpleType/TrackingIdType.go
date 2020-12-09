package SimpleType

type TrackingIdType string

const (
	TrackingIdTypeEXPRESS TrackingIdType = "EXPRESS"

	TrackingIdTypeFEDEX TrackingIdType = "FEDEX"

	TrackingIdTypeGROUND TrackingIdType = "GROUND"

	TrackingIdTypeUSPS TrackingIdType = "USPS"
)
