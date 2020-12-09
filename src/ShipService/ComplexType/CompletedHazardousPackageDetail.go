package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type CompletedHazardousPackageDetail struct {

	// A unique reference id that matches the package to a package configuration. This is populated if the client provided a package configuration for several packages that have the exact same dangerous goods content.
	ReferenceId string `xml:"ReferenceId,omitempty"`

	Accessibility *SimpleType.DangerousGoodsAccessibilityType `xml:"Accessibility,omitempty"`

	// When true indicates that the package can be transported only on a cargo aircraft.
	CargoAircraftOnly bool `xml:"CargoAircraftOnly,omitempty"`

	Regulation *SimpleType.HazardousCommodityRegulationType `xml:"Regulation,omitempty"`

	// Specifies the maximum radiation level from the package (measured in microSieverts per hour at a distance of one meter from the external surface of the package, divided by 10).
	RadioactiveTransportIndex float64 `xml:"RadioactiveTransportIndex,omitempty"`

	// Specifies the label that is to be put on a package containing radioactive material. The label type is determined in accordance with the Transportation of Dangerous Goods Act and indicates the type of radioactive material being handled as well as the relative risk.
	LabelType *SimpleType2.RadioactiveLabelType `xml:"LabelType,omitempty"`

	// Documents the kinds and quantities of all hazardous commodities in the current package.
	Containers []*ValidatedHazardousContainer `xml:"Containers,omitempty"`
}
