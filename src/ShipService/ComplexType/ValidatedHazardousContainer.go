package ComplexType

type ValidatedHazardousContainer struct {

	// Indicates that the quantity of the dangerous goods packaged is permissible for shipping. This is used to ensure that the dangerous goods commodities do not exceed the net quantity per package restrictions.
	QValue float64 `xml:"QValue,omitempty"`

	// Documents the kinds and quantities of all hazardous commodities in the current package.
	HazardousCommodities []*ValidatedHazardousCommodityContent `xml:"HazardousCommodities,omitempty"`
}
