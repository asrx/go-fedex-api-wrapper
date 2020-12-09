package ComplexType

import "time"

type CompletedTagDetail struct {
	ConfirmationNumber string `xml:"ConfirmationNumber,omitempty"`

	// As of June 2007, returned only for FedEx Express services.
	AccessTime *time.Duration `xml:"AccessTime,omitempty"`

	// As of June 2007, returned only for FedEx Express services.
	CutoffTime string `xml:"CutoffTime,omitempty"`

	// As of June 2007, returned only for FedEx Express services.
	Location string `xml:"Location,omitempty"`

	// As of June 2007, returned only for FedEx Express services.
	DeliveryCommitment string `xml:"DeliveryCommitment,omitempty"`

	// FEDEX INTERNAL USE ONLY: for use by INET.
	DispatchDate string `xml:"DispatchDate,omitempty"`
}
