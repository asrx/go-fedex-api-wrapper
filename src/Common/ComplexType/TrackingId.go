package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type TrackingId struct {
	TrackingIdType *SimpleType.TrackingIdType `xml:"TrackingIdType,omitempty"`

	FormId string `xml:"FormId,omitempty"`

	TrackingNumber string `xml:"TrackingNumber,omitempty"`
}
