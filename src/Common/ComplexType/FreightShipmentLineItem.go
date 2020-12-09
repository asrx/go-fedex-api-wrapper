package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type FreightShipmentLineItem struct {

	// A unique identifier assigned to this line item.
	Id string `xml:"Id,omitempty"`

	// Freight class for this line item.
	FreightClass *FreightClassType `xml:"FreightClass,omitempty"`

	// Specification of handling-unit packaging for this commodity or class line.
	Packaging *PhysicalPackagingType `xml:"Packaging,omitempty"`

	// Number of pieces for this commodity or class line.
	Pieces *uint `xml:"Pieces,omitempty"`

	// Customer-provided description for this commodity or class line.
	Description string `xml:"Description,omitempty"`

	// Weight for this commodity or class line.
	Weight *Weight `xml:"Weight,omitempty"`

	Dimensions *Dimensions `xml:"Dimensions,omitempty"`

	// Volume (cubic measure) for this commodity or class line.
	Volume *Volume `xml:"Volume,omitempty"`
}
