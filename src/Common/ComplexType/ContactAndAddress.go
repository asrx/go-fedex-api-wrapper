package ComplexType

type ContactAndAddress struct {
	Contact *Contact `xml:"Contact,omitempty"`

	Address *Address `xml:"Address,omitempty"`
}
