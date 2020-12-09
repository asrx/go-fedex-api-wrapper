package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type SmartPostShipmentProcessingOptionsRequested struct {
	Options []*SimpleType.SmartPostShipmentProcessingOptionType `xml:"Options,omitempty"`
}
