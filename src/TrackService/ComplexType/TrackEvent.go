package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackEvent struct {

	// The time this event occurred.
	Timestamp string `xml:"Timestamp,omitempty"`

	// Carrier's scan code. Pairs with EventDescription.
	EventType string `xml:"EventType,omitempty"`

	// Literal description that pairs with the EventType.
	EventDescription string `xml:"EventDescription,omitempty"`

	// Further defines the Scan Type code's specific type (e.g., DEX08 business closed). Pairs with StatusExceptionDescription.
	StatusExceptionCode string `xml:"StatusExceptionCode,omitempty"`

	// Literal description that pairs with the StatusExceptionCode.
	StatusExceptionDescription string `xml:"StatusExceptionDescription,omitempty"`

	// Address information of the station that is responsible for the scan.
	Address *ComplexType.Address `xml:"Address,omitempty"`

	// FedEx location ID where the scan took place. (Returned for CSR SL only.)
	StationId string `xml:"StationId,omitempty"`

	// Indicates where the arrival actually occurred.
	ArrivalLocation *SimpleType.ArrivalLocationType `xml:"ArrivalLocation,omitempty"`
}
