package ComplexType


type Measure struct {
	Quantity float64 `xml:"Quantity,omitempty"`

	Units string `xml:"Units,omitempty"`
}
