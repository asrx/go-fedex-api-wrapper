package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/RateService/SimpleType"

type FlatbedTrailerDetail struct {
	Options []*SimpleType.FlatbedTrailerOption `xml:"Options,omitempty"`
}
