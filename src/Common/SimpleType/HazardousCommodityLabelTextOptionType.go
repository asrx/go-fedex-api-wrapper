package SimpleType

// Specifies how the commodity is to be labeled.
type HazardousCommodityLabelTextOptionType string

const (
	HazardousCommodityLabelTextOptionTypeAPPEND HazardousCommodityLabelTextOptionType = "APPEND"

	HazardousCommodityLabelTextOptionTypeOVERRIDE HazardousCommodityLabelTextOptionType = "OVERRIDE"

	HazardousCommodityLabelTextOptionTypeSTANDARD HazardousCommodityLabelTextOptionType = "STANDARD"
)