package ComplexType

type CustomLabelGraphicEntry struct {
	Position *CustomLabelPosition `xml:"Position,omitempty"`

	// Printer-specific index of graphic image to be printed.
	PrinterGraphicId string `xml:"PrinterGraphicId,omitempty"`

	// Fully-qualified path and file name for graphic image to be printed.
	FileGraphicFullName string `xml:"FileGraphicFullName,omitempty"`
}
