package ComplexType

type ReturnAssociationDetail struct {

	// Specifies the tracking number of the master associated with the return shipment.
	TrackingNumber string `xml:"TrackingNumber,omitempty"`

	ShipDate string `xml:"ShipDate,omitempty"`
}
