package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type DangerousGoodsDetail struct {
	Regulation *SimpleType.HazardousCommodityRegulationType `xml:"Regulation,omitempty"`

	Accessibility *SimpleType.DangerousGoodsAccessibilityType `xml:"Accessibility,omitempty"`

	// Shipment is packaged/documented for movement ONLY on cargo aircraft.
	CargoAircraftOnly bool `xml:"CargoAircraftOnly,omitempty"`

	// Indicates which kinds of hazardous content are in the current package.
	Options []*SimpleType.HazardousCommodityOptionType `xml:"Options,omitempty"`

	// Indicates whether there is additional customer provided packaging enclosing the approved dangerous goods containers.
	PackingOption *SimpleType.DangerousGoodsPackingOptionType `xml:"PackingOption,omitempty"`

	// Identifies the configuration of this dangerous goods package. The common configuration is represented at the shipment level.
	ReferenceId string `xml:"ReferenceId,omitempty"`

	// Indicates one or more containers used to pack dangerous goods commodities.
	Containers []*DangerousGoodsContainer `xml:"Containers,omitempty"`

	// Description of the packaging of this commodity, suitable for use on OP-900 and OP-950 forms.
	Packaging *HazardousCommodityPackagingDetail `xml:"Packaging,omitempty"`

	// Name, title and place of the signatory for this shipment.
	Signatory *DangerousGoodsSignatory `xml:"Signatory,omitempty"`

	// Telephone number to use for contact in the event of an emergency.
	EmergencyContactNumber string `xml:"EmergencyContactNumber,omitempty"`

	// Offeror's name or contract number, per DOT regulation.
	Offeror string `xml:"Offeror,omitempty"`

	// Specifies the contact of the party responsible for handling the infectious substances, if any, in the dangerous goods shipment.
	InfectiousSubstanceResponsibleContact *Contact `xml:"InfectiousSubstanceResponsibleContact,omitempty"`

	// Specifies additional handling information for the current package.
	AdditionalHandling string `xml:"AdditionalHandling,omitempty"`

	// Specifies the radioactivity detail for the current package, if the package contains radioactive materials.
	RadioactivityDetail *RadioactivityDetail `xml:"RadioactivityDetail,omitempty"`
}
