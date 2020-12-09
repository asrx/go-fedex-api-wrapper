package SimpleType

// Relative to normal orientation for the printer.
type LabelRotationType string

const (
	LabelRotationTypeLEFT LabelRotationType = "LEFT"

	LabelRotationTypeNONE LabelRotationType = "NONE"

	LabelRotationTypeRIGHT LabelRotationType = "RIGHT"

	LabelRotationTypeUPSIDE_DOWN LabelRotationType = "UPSIDE_DOWN"
)