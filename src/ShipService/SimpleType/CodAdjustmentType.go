package SimpleType


// Specifies the type of adjustment was performed to the COD collection amount during rating.
type CodAdjustmentType string

const (
	CodAdjustmentTypeCHARGES_ADDED CodAdjustmentType = "CHARGES_ADDED"

	CodAdjustmentTypeNONE CodAdjustmentType = "NONE"
)

