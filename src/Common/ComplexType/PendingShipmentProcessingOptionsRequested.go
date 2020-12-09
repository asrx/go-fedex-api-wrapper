package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type PendingShipmentProcessingOptionsRequested struct {
	Options []*SimpleType.PendingShipmentProcessingOptionType `xml:"Options,omitempty"`
}
