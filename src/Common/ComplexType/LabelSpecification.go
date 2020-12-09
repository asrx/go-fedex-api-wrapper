package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type LabelSpecification struct {
	LabelFormatType *SimpleType.LabelFormatType `xml:"LabelFormatType,omitempty"`

	ImageType *SimpleType.ShippingDocumentImageType `xml:"ImageType,omitempty"`

	LabelStockType *SimpleType.LabelStockType `xml:"LabelStockType,omitempty"`

	LabelPrintingOrientation *SimpleType.LabelPrintingOrientationType `xml:"LabelPrintingOrientation,omitempty"`

	LabelRotation *SimpleType.LabelRotationType `xml:"LabelRotation,omitempty"`

	// Specifies the order in which the labels are requested to be returned
	LabelOrder *SimpleType.LabelOrderType `xml:"LabelOrder,omitempty"`

	PrintedLabelOrigin *ContactAndAddress `xml:"PrintedLabelOrigin,omitempty"`

	CustomerSpecifiedDetail *CustomerSpecifiedLabelDetail `xml:"CustomerSpecifiedDetail,omitempty"`
}
