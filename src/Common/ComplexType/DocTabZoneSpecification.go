package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DocTabZoneSpecification struct {
	ZoneNumber *uint `xml:"ZoneNumber,omitempty"`

	Header string `xml:"Header,omitempty"`

	DataField string `xml:"DataField,omitempty"`

	LiteralValue string `xml:"LiteralValue,omitempty"`

	Justification *SimpleType.DocTabZoneJustificationType `xml:"Justification,omitempty"`
}
