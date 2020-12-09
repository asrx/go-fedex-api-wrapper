package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DocumentFormatOptionsRequested struct {
	Options []*SimpleType.DocumentFormatOptionType `xml:"Options,omitempty"`
}
