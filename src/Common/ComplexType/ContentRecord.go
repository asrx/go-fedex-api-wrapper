package ComplexType

type ContentRecord struct {
	PartNumber string `xml:"PartNumber,omitempty"`

	ItemNumber string `xml:"ItemNumber,omitempty"`

	ReceivedQuantity *uint `xml:"ReceivedQuantity,omitempty"`

	Description string `xml:"Description,omitempty"`
}
