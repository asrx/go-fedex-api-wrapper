package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type CompletedCodDetail struct {
	CollectionAmount *ComplexType.Money `xml:"CollectionAmount,omitempty"`

	AdjustmentType *SimpleType.CodAdjustmentType `xml:"AdjustmentType,omitempty"`
}
