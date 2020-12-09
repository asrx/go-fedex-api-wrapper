package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type BatteryClassificationDetail struct {

	// Describes the material composition of the battery or cell.
	Material *SimpleType.BatteryMaterialType `xml:"Material,omitempty"`

	// Describes the packing arrangement of the battery or cell with respect to other items within the same package.
	Packing *SimpleType.BatteryPackingType `xml:"Packing,omitempty"`

	// A regulation specific classification for the battery or cell.
	RegulatorySubType *SimpleType.BatteryRegulatorySubType `xml:"RegulatorySubType,omitempty"`
}
