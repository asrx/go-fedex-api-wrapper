package SimpleType

// Specifies the measure of quantity to be validated against a prescribed limit.
type HazardousCommodityQuantityType string

const (
	HazardousCommodityQuantityTypeGROSS HazardousCommodityQuantityType = "GROSS"

	HazardousCommodityQuantityTypeNET HazardousCommodityQuantityType = "NET"
)