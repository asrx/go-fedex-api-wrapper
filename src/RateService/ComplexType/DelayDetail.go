package ComplexType

import (
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/RateService/SimpleType"
)

type DelayDetail struct {

	// The date of the delay
	Date string `xml:"Date,omitempty"`

	DayOfWeek *SimpleType2.DayOfWeekType `xml:"DayOfWeek,omitempty"`

	// The attribute of the shipment that caused the delay(e.g. Country, City, LocationId, Zip, service area, special handling )
	Level *SimpleType.DelayLevelType `xml:"Level,omitempty"`

	// The point where the delay is occurring (e.g. Origin, Destination, Broker location)
	Point *SimpleType.DelayPointType `xml:"Point,omitempty"`

	// The reason for the delay (e.g. holiday, weekend, etc.).
	Type *SimpleType.CommitmentDelayType `xml:"Type,omitempty"`

	// The name of the holiday in that country that is causing the delay.
	Description string `xml:"Description,omitempty"`
}
