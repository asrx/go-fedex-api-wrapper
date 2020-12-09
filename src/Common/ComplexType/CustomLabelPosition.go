package ComplexType

type CustomLabelPosition struct {

	// Horizontal position, relative to left edge of custom area.
	X *uint `xml:"X,omitempty"`

	// Vertical position, relative to top edge of custom area.
	Y int32 `xml:"Y,omitempty"`
}
