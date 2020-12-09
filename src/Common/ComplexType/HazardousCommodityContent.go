package ComplexType

type HazardousCommodityContent struct {

	// Identifies and describes an individual hazardous commodity.
	Description *HazardousCommodityDescription `xml:"Description,omitempty"`

	// Specifies the amount of the commodity in alternate units.
	Quantity *HazardousCommodityQuantityDetail `xml:"Quantity,omitempty"`

	// This describes the inner receptacle details for a hazardous commodity within the dangerous goods container.
	InnerReceptacles []*HazardousCommodityInnerReceptacleDetail `xml:"InnerReceptacles,omitempty"`

	// Customer-provided specifications for handling individual commodities.
	Options *HazardousCommodityOptionDetail `xml:"Options,omitempty"`

	// Specifies the details of any radio active materials within the commodity.
	RadionuclideDetail *RadionuclideDetail `xml:"RadionuclideDetail,omitempty"`

	// The total mass of the contained explosive substances, without the mass of any casings, bullets, shells, etc.
	NetExplosiveDetail *NetExplosiveDetail `xml:"NetExplosiveDetail,omitempty"`
}
