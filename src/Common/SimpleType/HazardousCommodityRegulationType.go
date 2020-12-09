package SimpleType

// Identifies the source of regulation for hazardous commodity data.
type HazardousCommodityRegulationType string

const (
	HazardousCommodityRegulationTypeADR HazardousCommodityRegulationType = "ADR"

	HazardousCommodityRegulationTypeDOT HazardousCommodityRegulationType = "DOT"

	HazardousCommodityRegulationTypeIATA HazardousCommodityRegulationType = "IATA"

	HazardousCommodityRegulationTypeORMD HazardousCommodityRegulationType = "ORMD"
)