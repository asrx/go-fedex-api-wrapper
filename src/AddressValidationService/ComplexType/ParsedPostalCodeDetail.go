package ComplexType

type ParsedPostalCodeDetail struct {
	Base string `xml:"Base,omitempty"`

	AddOn string `xml:"AddOn,omitempty"`

	DeliveryPoint string `xml:"DeliveryPoint,omitempty"`
}
