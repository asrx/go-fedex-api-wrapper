package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomLabelDetail struct {
	CoordinateUnits *SimpleType.CustomLabelCoordinateUnits `xml:"CoordinateUnits,omitempty"`

	TextEntries []*CustomLabelTextEntry `xml:"TextEntries,omitempty"`

	GraphicEntries []*CustomLabelGraphicEntry `xml:"GraphicEntries,omitempty"`

	BoxEntries []*CustomLabelBoxEntry `xml:"BoxEntries,omitempty"`

	TextBoxEntries []*CustomLabelTextBoxEntry `xml:"TextBoxEntries,omitempty"`

	BarcodeEntries []*CustomLabelBarcodeEntry `xml:"BarcodeEntries,omitempty"`
}
