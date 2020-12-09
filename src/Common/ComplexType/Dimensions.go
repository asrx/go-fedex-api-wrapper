package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Dimensions struct {
	Length *uint `xml:"Length,omitempty"`

	Width *uint `xml:"Width,omitempty"`

	Height *uint `xml:"Height,omitempty"`

	Units *SimpleType.LinearUnits `xml:"Units,omitempty"`
}
