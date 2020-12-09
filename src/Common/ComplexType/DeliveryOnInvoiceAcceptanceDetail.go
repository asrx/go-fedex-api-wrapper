package ComplexType

type DeliveryOnInvoiceAcceptanceDetail struct {
	Recipient *Party `xml:"Recipient,omitempty"`

	// Specifies the tracking id for the return, if preassigned.
	TrackingId *TrackingId `xml:"TrackingId,omitempty"`
}
