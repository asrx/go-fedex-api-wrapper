package ComplexType


type AppointmentDetail struct {
	Date string `xml:"Date,omitempty"`

	// Different appointment time windows on the date specified.
	WindowDetails []*AppointmentTimeDetail `xml:"WindowDetails,omitempty"`
}
