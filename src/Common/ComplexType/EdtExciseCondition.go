package ComplexType

type EdtExciseCondition struct {
	Category string `xml:"Category,omitempty"`

	// Customer-declared value, with data type and legal values depending on excise condition, used in defining the taxable value of the item.
	Value string `xml:"Value,omitempty"`
}
