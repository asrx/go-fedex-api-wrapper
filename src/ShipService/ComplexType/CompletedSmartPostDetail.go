package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CompletedSmartPostDetail struct {

	// Identifies the carrier that will pick up the SmartPost shipment.
	PickUpCarrier *SimpleType.CarrierCodeType `xml:"PickUpCarrier,omitempty"`

	// Indicates whether the shipment is deemed to be machineable, based on dimensions, weight, and packaging.
	Machinable bool `xml:"Machinable,omitempty"`
}
