package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomLabelTextBoxEntry struct {
	TopLeftCorner *CustomLabelPosition `xml:"TopLeftCorner,omitempty"`

	BottomRightCorner *CustomLabelPosition `xml:"BottomRightCorner,omitempty"`

	Position *CustomLabelPosition `xml:"Position,omitempty"`

	Format string `xml:"Format,omitempty"`

	DataFields []string `xml:"DataFields,omitempty"`

	// Printer-specific font name for use with thermal printer labels.
	ThermalFontId string `xml:"ThermalFontId,omitempty"`

	// Generic font name for use with plain paper labels.
	FontName string `xml:"FontName,omitempty"`

	// Generic font size for use with plain paper labels.
	FontSize *uint `xml:"FontSize,omitempty"`

	Rotation *SimpleType.RotationType `xml:"Rotation,omitempty"`
}
