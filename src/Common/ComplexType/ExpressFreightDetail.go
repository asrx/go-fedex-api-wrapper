package ComplexType

type ExpressFreightDetail struct {
	PackingListEnclosed bool `xml:"PackingListEnclosed,omitempty"`

	ShippersLoadAndCount *uint `xml:"ShippersLoadAndCount,omitempty"`

	BookingConfirmationNumber string `xml:"BookingConfirmationNumber,omitempty"`

	ReferenceLabelRequested bool `xml:"ReferenceLabelRequested,omitempty"`

	BeforeDeliveryContact *ExpressFreightDetailContact `xml:"BeforeDeliveryContact,omitempty"`

	UndeliverableContact *ExpressFreightDetailContact `xml:"UndeliverableContact,omitempty"`
}
