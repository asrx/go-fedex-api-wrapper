package SimpleType

// Indicates which kind of hazardous content is being reported.
type HazardousCommodityOptionType string

const (
	HazardousCommodityOptionTypeBATTERY HazardousCommodityOptionType = "BATTERY"

	HazardousCommodityOptionTypeHAZARDOUS_MATERIALS HazardousCommodityOptionType = "HAZARDOUS_MATERIALS"

	HazardousCommodityOptionTypeLIMITED_QUANTITIES_COMMODITIES HazardousCommodityOptionType = "LIMITED_QUANTITIES_COMMODITIES"

	HazardousCommodityOptionTypeORM_D HazardousCommodityOptionType = "ORM_D"

	HazardousCommodityOptionTypeREPORTABLE_QUANTITIES HazardousCommodityOptionType = "REPORTABLE_QUANTITIES"

	HazardousCommodityOptionTypeSMALL_QUANTITY_EXCEPTION HazardousCommodityOptionType = "SMALL_QUANTITY_EXCEPTION"
)