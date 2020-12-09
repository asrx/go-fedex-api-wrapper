package ComplexType

type DangerousGoodsSignatory struct {
	ContactName string `xml:"ContactName,omitempty"`

	Title string `xml:"Title,omitempty"`

	// Indicates the place where the form is signed.
	Place string `xml:"Place,omitempty"`
}
