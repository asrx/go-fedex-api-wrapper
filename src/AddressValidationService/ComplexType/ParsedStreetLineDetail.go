package ComplexType

type ParsedStreetLineDetail struct {
	HouseNumber string `xml:"HouseNumber,omitempty"`

	PreStreetType string `xml:"PreStreetType,omitempty"`

	LeadingDirectional string `xml:"LeadingDirectional,omitempty"`

	StreetName string `xml:"StreetName,omitempty"`

	StreetSuffix string `xml:"StreetSuffix,omitempty"`

	TrailingDirectional string `xml:"TrailingDirectional,omitempty"`

	UnitLabel string `xml:"UnitLabel,omitempty"`

	UnitNumber string `xml:"UnitNumber,omitempty"`

	RuralRoute string `xml:"RuralRoute,omitempty"`

	POBox string `xml:"POBox,omitempty"`

	Building string `xml:"Building,omitempty"`

	Organization string `xml:"Organization,omitempty"`
}
