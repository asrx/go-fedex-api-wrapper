package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
)

type TrackStatusDetail struct {
	CreationTime string `xml:"CreationTime,omitempty"`

	Code string `xml:"Code,omitempty"`

	Description string `xml:"Description,omitempty"`

	Location *ComplexType.Address `xml:"Location,omitempty"`

	AncillaryDetails []*TrackStatusAncillaryDetail `xml:"AncillaryDetails,omitempty"`
}
