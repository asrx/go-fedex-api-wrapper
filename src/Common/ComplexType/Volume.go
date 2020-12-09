package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Volume struct {
	Units *SimpleType.VolumeUnits `xml:"Units,omitempty"`

	Value float64 `xml:"Value,omitempty"`
}
