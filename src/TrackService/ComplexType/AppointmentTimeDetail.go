package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type AppointmentTimeDetail struct {

	// The description that FedEx Ground uses for the appointment window being specified.
	Type *SimpleType.AppointmentWindowType `xml:"Type,omitempty"`

	// Specifies the window of time for an appointment.
	Window *LocalTimeRange `xml:"Window,omitempty"`

	Description string `xml:"Description,omitempty"`
}
