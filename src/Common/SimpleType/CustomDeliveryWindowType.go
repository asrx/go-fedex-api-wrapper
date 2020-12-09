package SimpleType


type CustomDeliveryWindowType string

const (
	CustomDeliveryWindowTypeAFTER CustomDeliveryWindowType = "AFTER"

	CustomDeliveryWindowTypeBEFORE CustomDeliveryWindowType = "BEFORE"

	CustomDeliveryWindowTypeBETWEEN CustomDeliveryWindowType = "BETWEEN"

	CustomDeliveryWindowTypeON CustomDeliveryWindowType = "ON"
)