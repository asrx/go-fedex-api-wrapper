package SimpleType

type RateDiscountType string

const (
	RateDiscountTypeBONUS RateDiscountType = "BONUS"

	RateDiscountTypeCOUPON RateDiscountType = "COUPON"

	RateDiscountTypeEARNED RateDiscountType = "EARNED"

	RateDiscountTypeOTHER RateDiscountType = "OTHER"

	RateDiscountTypeVOLUME RateDiscountType = "VOLUME"
)
