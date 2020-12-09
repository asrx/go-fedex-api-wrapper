package SimpleType

type MinimumChargeType string

const (
	MinimumChargeTypeCUSTOMER MinimumChargeType = "CUSTOMER"

	MinimumChargeTypeCUSTOMER_FREIGHT_WEIGHT MinimumChargeType = "CUSTOMER_FREIGHT_WEIGHT"

	MinimumChargeTypeEARNED_DISCOUNT MinimumChargeType = "EARNED_DISCOUNT"

	MinimumChargeTypeMIXED MinimumChargeType = "MIXED"

	MinimumChargeTypeRATE_SCALE MinimumChargeType = "RATE_SCALE"
)