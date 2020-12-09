package ComplexType


type AddressAttribute struct {

	// Specifies the key for the address attribute.
	Name string `xml:"Name,omitempty"`

	// The value for the key for the address attribute.
	Value string `xml:"Value,omitempty"`
}
