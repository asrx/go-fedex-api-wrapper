package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackingDateOrTimestamp struct {
	Type *SimpleType.TrackingDateOrTimestampType `xml:"Type,omitempty"`

	DateOrTimestamp string `xml:"DateOrTimestamp,omitempty"`
}
