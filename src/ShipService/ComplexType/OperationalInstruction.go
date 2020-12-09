package ComplexType

type OperationalInstruction struct {

	// Position of operational instruction element.
	Number int32 `xml:"Number,omitempty"`

	// Content corresponding to the operational instruction.
	Content string `xml:"Content,omitempty"`
}
