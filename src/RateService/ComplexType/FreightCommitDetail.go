package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type FreightCommitDetail struct {

	// Information about the origin Freight Service Center.
	OriginDetail *FreightServiceCenterDetail `xml:"OriginDetail,omitempty"`

	// Information about the destination Freight Service Center.
	DestinationDetail *FreightServiceCenterDetail `xml:"DestinationDetail,omitempty"`

	// The distance between the origin and destination FreightService Centers
	TotalDistance *ComplexType.Distance `xml:"TotalDistance,omitempty"`
}
