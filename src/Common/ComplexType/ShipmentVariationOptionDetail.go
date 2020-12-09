package ComplexType

type ShipmentVariationOptionDetail struct {

	// Specifies the name or the key for the shipment variation.
	Id string `xml:"Id,omitempty"`

	// The values that are valid for the specified shipment variation in the context of the current shipment.
	Values []string `xml:"Values,omitempty"`
}
