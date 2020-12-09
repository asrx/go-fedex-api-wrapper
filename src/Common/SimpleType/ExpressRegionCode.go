package SimpleType

// Indicates a FedEx Express operating region.
type ExpressRegionCode string

const (
	ExpressRegionCodeAPAC ExpressRegionCode = "APAC"

	ExpressRegionCodeCA ExpressRegionCode = "CA"

	ExpressRegionCodeEMEA ExpressRegionCode = "EMEA"

	ExpressRegionCodeLAC ExpressRegionCode = "LAC"

	ExpressRegionCodeUS ExpressRegionCode = "US"
)