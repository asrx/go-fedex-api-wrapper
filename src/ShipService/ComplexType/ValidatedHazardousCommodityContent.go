package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type ValidatedHazardousCommodityContent struct {

	// Identifies and describes an individual hazardous commodity.
	Description *ValidatedHazardousCommodityDescription `xml:"Description,omitempty"`

	// Specifies the amount of the commodity in alternate units.
	Quantity *ComplexType.HazardousCommodityQuantityDetail `xml:"Quantity,omitempty"`

	// The mass points are a calculation used by ADR regulations for measuring the risk of a particular hazardous commodity.
	MassPoints float64 `xml:"MassPoints,omitempty"`

	// Customer-provided specifications for handling individual commodities.
	Options *ComplexType.HazardousCommodityOptionDetail `xml:"Options,omitempty"`

	// The total mass of the contained explosive substances, without the mass of any casings, bullets, shells, etc.
	NetExplosiveDetail *ComplexType.NetExplosiveDetail `xml:"NetExplosiveDetail,omitempty"`
}
