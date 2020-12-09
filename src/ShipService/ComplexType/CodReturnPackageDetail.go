package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type CodReturnPackageDetail struct {
	CollectionAmount *ComplexType.Money `xml:"CollectionAmount,omitempty"`

	AdjustmentType *SimpleType.CodAdjustmentType `xml:"AdjustmentType,omitempty"`

	Electronic bool `xml:"Electronic,omitempty"`

	Barcodes *PackageBarcodes `xml:"Barcodes,omitempty"`

	Label *ShippingDocument `xml:"Label,omitempty"`
}
