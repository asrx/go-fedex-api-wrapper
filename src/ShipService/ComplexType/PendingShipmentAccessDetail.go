package ComplexType

type PendingShipmentAccessDetail struct {
	AccessorDetails []*PendingShipmentAccessorDetail `xml:"AccessorDetails,omitempty"`
}
