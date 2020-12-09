package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type NetExplosiveDetail struct {
	Type *SimpleType.NetExplosiveClassificationType `xml:"Type,omitempty"`

	Amount float64 `xml:"Amount,omitempty"`

	Units string `xml:"Units,omitempty"`
}
