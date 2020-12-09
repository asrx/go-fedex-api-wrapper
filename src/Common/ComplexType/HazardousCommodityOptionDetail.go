package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type HazardousCommodityOptionDetail struct {

	// Specifies how the customer wishes the label text to be handled for this commodity in this package.
	LabelTextOption *SimpleType.HazardousCommodityLabelTextOptionType `xml:"LabelTextOption,omitempty"`

	// Text used in labeling the commodity under control of the labelTextOption field.
	CustomerSuppliedLabelText string `xml:"CustomerSuppliedLabelText,omitempty"`
}
