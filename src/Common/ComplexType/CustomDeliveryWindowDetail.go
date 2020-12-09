package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CustomDeliveryWindowDetail struct {

	// Indicates the type of custom delivery being requested.
	Type *SimpleType.CustomDeliveryWindowType `xml:"Type,omitempty"`

	// Time by which delivery is requested.
	RequestTime string `xml:"RequestTime,omitempty"`

	// Range of dates for custom delivery request; only used if type is BETWEEN.
	RequestRange *DateRange `xml:"RequestRange,omitempty"`

	// Date for custom delivery request; only used for types of ON, BETWEEN, or AFTER.
	RequestDate string `xml:"RequestDate,omitempty"`
}