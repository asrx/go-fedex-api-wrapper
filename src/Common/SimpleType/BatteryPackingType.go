package SimpleType


// Describes the packing arrangement of a battery or cell with respect to other items within the same package.
type BatteryPackingType string

const (
	BatteryPackingTypeCONTAINED_IN_EQUIPMENT BatteryPackingType = "CONTAINED_IN_EQUIPMENT"

	BatteryPackingTypePACKED_WITH_EQUIPMENT BatteryPackingType = "PACKED_WITH_EQUIPMENT"
)
