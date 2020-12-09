package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type HazardousCommodityDescription struct {

	// Regulatory identifier for a commodity (e.g. "UN ID" value).
	Id string `xml:"Id,omitempty"`

	// In conjunction with the regulatory identifier, this field uniquely identifies a specific hazardous materials commodity.
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	PackingGroup *SimpleType.HazardousCommodityPackingGroupType `xml:"PackingGroup,omitempty"`

	PackingDetails *HazardousCommodityPackingDetail `xml:"PackingDetails,omitempty"`

	ReportableQuantity bool `xml:"ReportableQuantity,omitempty"`

	ProperShippingName string `xml:"ProperShippingName,omitempty"`

	TechnicalName string `xml:"TechnicalName,omitempty"`

	Percentage float64 `xml:"Percentage,omitempty"`

	HazardClass string `xml:"HazardClass,omitempty"`

	SubsidiaryClasses []string `xml:"SubsidiaryClasses,omitempty"`

	LabelText string `xml:"LabelText,omitempty"`

	// Indicates any special processing options to be applied to the description of the dangerous goods commodity.
	ProcessingOptions []*SimpleType.HazardousCommodityDescriptionProcessingOptionType `xml:"ProcessingOptions,omitempty"`

	// Information related to quantity limitations and operator or state variations as may be applicable to the dangerous goods commodity.
	Authorization string `xml:"Authorization,omitempty"`
}
