package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DangerousGoodsContainer struct {

	// Indicates whether there are additional inner receptacles within this container.
	PackingType *SimpleType.HazardousContainerPackingType `xml:"PackingType,omitempty"`

	// Indicates the type of this dangerous goods container, as specified by the IATA packing instructions. For example, steel cylinder, fiberboard box, plastic jerrican and steel drum.
	ContainerType string `xml:"ContainerType,omitempty"`

	// Indicates the packaging type of the container used to package the radioactive materials.
	RadioactiveContainerClass *SimpleType.RadioactiveContainerClassType `xml:"RadioactiveContainerClass,omitempty"`

	// Indicates the number of occurrences of this container with identical dangerous goods configuration.
	NumberOfContainers *uint `xml:"NumberOfContainers,omitempty"`

	// Documents the kinds and quantities of all hazardous commodities in the current container.
	HazardousCommodities []*HazardousCommodityContent `xml:"HazardousCommodities,omitempty"`
}
