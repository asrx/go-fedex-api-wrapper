package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type TrackServiceDescriptionDetail struct {
	Type *SimpleType.ServiceType `xml:"Type,omitempty"`

	Description string `xml:"Description,omitempty"`

	// Specifies a shorter description for the service that is calculated per the service code.
	ShortDescription string `xml:"ShortDescription,omitempty"`
}
