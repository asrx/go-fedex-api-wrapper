package ComplexType

type FreightRateNotation struct {

	// Unique identifier for notation.
	Code string `xml:"Code,omitempty"`

	// Human-readable explanation of notation.
	Description string `xml:"Description,omitempty"`
}
