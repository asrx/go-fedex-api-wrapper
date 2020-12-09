package SimpleType

// Indicates the reason that a dim divisor value was chose.
type RateDimensionalDivisorType string

const (
	RateDimensionalDivisorTypeCOUNTRY RateDimensionalDivisorType = "COUNTRY"

	RateDimensionalDivisorTypeCUSTOMER RateDimensionalDivisorType = "CUSTOMER"

	RateDimensionalDivisorTypeOTHER RateDimensionalDivisorType = "OTHER"

	RateDimensionalDivisorTypePRODUCT RateDimensionalDivisorType = "PRODUCT"

	RateDimensionalDivisorTypeWAIVED RateDimensionalDivisorType = "WAIVED"
)
