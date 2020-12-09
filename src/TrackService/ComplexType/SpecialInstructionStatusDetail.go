package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type SpecialInstructionStatusDetail struct {

	// Specifies the status of the track special instructions requested.
	Status *SimpleType.SpecialInstructionsStatusCode `xml:"Status,omitempty"`

	// Time when the status was changed.
	StatusCreateTime string `xml:"StatusCreateTime,omitempty"`
}
