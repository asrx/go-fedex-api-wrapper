package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CodAddTransportationChargesDetail struct {

	// Select the type of rate from which the element is to be selected.
	RateTypeBasis *SimpleType.RateTypeBasisType `xml:"RateTypeBasis,omitempty"`

	ChargeBasis *SimpleType.CodAddTransportationChargeBasisType `xml:"ChargeBasis,omitempty"`

	ChargeBasisLevel *SimpleType.ChargeBasisLevelType `xml:"ChargeBasisLevel,omitempty"`
}
