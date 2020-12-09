package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type EmailOptionsRequested struct {
	Options []*SimpleType.EmailOptionType `xml:"Options,omitempty"`
}
