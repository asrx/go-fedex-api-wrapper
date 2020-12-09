package ComplexType

type DocumentLineItem struct {
	Values []*KeyValueDetail `xml:"Values,omitempty"`
}
