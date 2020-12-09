package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type PieceCountVerificationDetail struct {
	CountLocationType *SimpleType.PieceCountLocationType `xml:"CountLocationType,omitempty"`

	Count *uint `xml:"Count,omitempty"`

	Description string `xml:"Description,omitempty"`
}