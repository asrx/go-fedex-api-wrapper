package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShipmentDryIceProcessingOptionsRequested struct {
	Options []*SimpleType.ShipmentDryIceProcessingOptionType `xml:"Options,omitempty"`
}
