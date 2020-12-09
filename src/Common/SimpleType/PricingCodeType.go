package SimpleType


type PricingCodeType string

const (
	PricingCodeTypeACTUAL PricingCodeType = "ACTUAL"

	PricingCodeTypeALTERNATE PricingCodeType = "ALTERNATE"

	PricingCodeTypeBASE PricingCodeType = "BASE"

	PricingCodeTypeHUNDREDWEIGHT PricingCodeType = "HUNDREDWEIGHT"

	PricingCodeTypeHUNDREDWEIGHT_ALTERNATE PricingCodeType = "HUNDREDWEIGHT_ALTERNATE"

	PricingCodeTypeINTERNATIONAL_DISTRIBUTION PricingCodeType = "INTERNATIONAL_DISTRIBUTION"

	PricingCodeTypeINTERNATIONAL_ECONOMY_SERVICE PricingCodeType = "INTERNATIONAL_ECONOMY_SERVICE"

	PricingCodeTypeLTL_FREIGHT PricingCodeType = "LTL_FREIGHT"

	PricingCodeTypePACKAGE PricingCodeType = "PACKAGE"

	PricingCodeTypeSHIPMENT PricingCodeType = "SHIPMENT"

	PricingCodeTypeSHIPMENT_FIVE_POUND_OPTIONAL PricingCodeType = "SHIPMENT_FIVE_POUND_OPTIONAL"

	PricingCodeTypeSHIPMENT_OPTIONAL PricingCodeType = "SHIPMENT_OPTIONAL"

	PricingCodeTypeSPECIAL PricingCodeType = "SPECIAL"
)