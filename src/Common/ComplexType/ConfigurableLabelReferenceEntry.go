package ComplexType

type ConfigurableLabelReferenceEntry struct {
	ZoneNumber *uint `xml:"ZoneNumber,omitempty"`

	Header string `xml:"Header,omitempty"`

	DataField string `xml:"DataField,omitempty"`

	LiteralValue string `xml:"LiteralValue,omitempty"`
}
