package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type Commodity struct {

	// Value used to identify a commodity description; must be unique within the containing shipment.
	CommodityId string `xml:"CommodityId,omitempty"`

	// FedEx internal commodity identifier
	Name string `xml:"Name,omitempty"`

	NumberOfPieces *uint `xml:"NumberOfPieces,omitempty"`

	// A free-form description of the commodity, which could be used for customs clearance documentation.
	Description string `xml:"Description,omitempty"`

	Purpose *SimpleType.CommodityPurposeType `xml:"Purpose,omitempty"`

	CountryOfManufacture string `xml:"CountryOfManufacture,omitempty"`

	HarmonizedCode string `xml:"HarmonizedCode,omitempty"`

	Weight *Weight `xml:"Weight,omitempty"`

	Quantity float64 `xml:"Quantity,omitempty"`

	QuantityUnits string `xml:"QuantityUnits,omitempty"`

	// Contains only additional quantitative information other than weight and quantity to calculate duties and taxes.
	AdditionalMeasures []*Measure `xml:"AdditionalMeasures,omitempty"`

	UnitPrice *Money `xml:"UnitPrice,omitempty"`

	CustomsValue *Money `xml:"CustomsValue,omitempty"`

	// Defines additional characteristic of commodity used to calculate duties and taxes
	ExciseConditions []*EdtExciseCondition `xml:"ExciseConditions,omitempty"`

	ExportLicenseNumber string `xml:"ExportLicenseNumber,omitempty"`

	ExportLicenseExpirationDate string `xml:"ExportLicenseExpirationDate,omitempty"`

	CIMarksAndNumbers string `xml:"CIMarksAndNumbers,omitempty"`

	PartNumber string `xml:"PartNumber,omitempty"`

	// All data required for this commodity in NAFTA Certificate of Origin.
	NaftaDetail *NaftaCommodityDetail `xml:"NaftaDetail,omitempty"`
}

