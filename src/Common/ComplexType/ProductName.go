package ComplexType

type ProductName struct {

	// The type of branded, translated, and/or localized name to which this value refers.
	Type string `xml:"Type,omitempty"`

	// The character encoding used to represent this product name. For example, UTF-8.
	Encoding string `xml:"Encoding,omitempty"`

	Value string `xml:"Value,omitempty"`
}
