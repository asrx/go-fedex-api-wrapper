package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Weight struct {

	// Identifies the unit of measure associated with a weight value.
	Units *SimpleType.WeightUnits `xml:"Units,omitempty"`

	// Identifies the weight value of a package/shipment.
	Value float64 `xml:"Value,omitempty"`
}
