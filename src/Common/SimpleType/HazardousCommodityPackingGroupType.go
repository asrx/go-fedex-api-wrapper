package SimpleType

// Identifies DOT packing group for a hazardous commodity.
type HazardousCommodityPackingGroupType string

const (
	HazardousCommodityPackingGroupTypeDEFAULT HazardousCommodityPackingGroupType = "DEFAULT"

	HazardousCommodityPackingGroupTypeI HazardousCommodityPackingGroupType = "I"

	HazardousCommodityPackingGroupTypeII HazardousCommodityPackingGroupType = "II"

	HazardousCommodityPackingGroupTypeIII HazardousCommodityPackingGroupType = "III"
)
