package ComplexType

type CustomLabelBoxEntry struct {
	TopLeftCorner *CustomLabelPosition `xml:"TopLeftCorner,omitempty"`

	BottomRightCorner *CustomLabelPosition `xml:"BottomRightCorner,omitempty"`
}
