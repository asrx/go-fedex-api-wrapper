package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
)

type ValidatedHazardousCommodityDescription struct {

	// Regulatory identifier for a commodity (e.g. "UN ID" value).
	Id string `xml:"Id,omitempty"`

	// In conjunction with the regulatory identifier, this field uniquely identifies a specific hazardous materials commodity.
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	PackingGroup *SimpleType.HazardousCommodityPackingGroupType `xml:"PackingGroup,omitempty"`

	PackingInstructions string `xml:"PackingInstructions,omitempty"`

	ProperShippingName string `xml:"ProperShippingName,omitempty"`

	// Fully-expanded descriptive text for a hazardous commodity.
	ProperShippingNameAndDescription string `xml:"ProperShippingNameAndDescription,omitempty"`

	TechnicalName string `xml:"TechnicalName,omitempty"`

	HazardClass string `xml:"HazardClass,omitempty"`

	SubsidiaryClasses []string `xml:"SubsidiaryClasses,omitempty"`

	// Coded indications for special requirements or constraints.
	Symbols string `xml:"Symbols,omitempty"`

	TunnelRestrictionCode string `xml:"TunnelRestrictionCode,omitempty"`

	SpecialProvisions string `xml:"SpecialProvisions,omitempty"`

	Attributes []*SimpleType2.HazardousCommodityAttributeType `xml:"Attributes,omitempty"`

	Authorization string `xml:"Authorization,omitempty"`

	LabelText string `xml:"LabelText,omitempty"`
}
