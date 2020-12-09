package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackReturnDetail struct {
	MovementStatus *SimpleType.TrackReturnMovementStatusType `xml:"MovementStatus,omitempty"`

	LabelType *SimpleType.TrackReturnLabelType `xml:"LabelType,omitempty"`

	Description string `xml:"Description,omitempty"`

	AuthorizationName string `xml:"AuthorizationName,omitempty"`
}
