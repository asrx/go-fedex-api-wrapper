package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type HazardousCommodityQuantityDetail struct {

	// Number of units of the type below.
	Amount float64 `xml:"Amount,omitempty"`

	// Units by which the hazardous commodity is measured. For IATA commodity, the units values are restricted based on regulation type.
	Units string `xml:"Units,omitempty"`

	// Specifies which measure of quantity is to be validated.
	QuantityType *SimpleType.HazardousCommodityQuantityType `xml:"QuantityType,omitempty"`
}
