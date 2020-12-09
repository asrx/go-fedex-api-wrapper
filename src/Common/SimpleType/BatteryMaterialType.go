package SimpleType


// Describes the material composition of a battery or cell.
type BatteryMaterialType string

const (
	BatteryMaterialTypeLITHIUM_ION BatteryMaterialType = "LITHIUM_ION"

	BatteryMaterialTypeLITHIUM_METAL BatteryMaterialType = "LITHIUM_METAL"
)