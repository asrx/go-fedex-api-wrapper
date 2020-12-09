package ComplexType


type Money struct {
	Currency string `xml:"Currency,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`
}
