package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackSpecialInstruction struct {
	Description string `xml:"Description,omitempty"`

	DeliveryOption *SimpleType.TrackDeliveryOptionType `xml:"DeliveryOption,omitempty"`

	// Specifies the status and status update time of the track special instructions.
	StatusDetail *SpecialInstructionStatusDetail `xml:"StatusDetail,omitempty"`

	// Specifies the estimated delivery time that was originally estimated when the shipment was shipped.
	OriginalEstimatedDeliveryTimestamp string `xml:"OriginalEstimatedDeliveryTimestamp,omitempty"`

	// Specifies the time the customer requested a change to the shipment.
	OriginalRequestTime string `xml:"OriginalRequestTime,omitempty"`

	// The requested appointment time for delivery.
	RequestedAppointmentTime *AppointmentDetail `xml:"RequestedAppointmentTime,omitempty"`
}
