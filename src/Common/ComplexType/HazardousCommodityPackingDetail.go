package ComplexType

type HazardousCommodityPackingDetail struct {
	CargoAircraftOnly bool `xml:"CargoAircraftOnly,omitempty"`

	// Coded specification for how commodity is to be packed.
	PackingInstructions string `xml:"PackingInstructions,omitempty"`
}
