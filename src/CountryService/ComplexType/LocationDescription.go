package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type LocationDescription struct {
	LocationId string `xml:"LocationId,omitempty"`

	LocationNumber int32 `xml:"LocationNumber,omitempty"`

	CountryCode string `xml:"CountryCode,omitempty"`

	StateOrProvinceCode string `xml:"StateOrProvinceCode,omitempty"`

	PostalCode string `xml:"PostalCode,omitempty"`

	ServiceArea string `xml:"ServiceArea,omitempty"`

	AirportId string `xml:"AirportId,omitempty"`

	// Package special services prohibited for this location for operational reasons. (Other package special services may or may not be available per shipment for other reasons.)
	RestrictedPackageSpecialServices []*SimpleType.PackageSpecialServiceType `xml:"RestrictedPackageSpecialServices,omitempty"`

	// Shipment special services prohibited for this location for operational reasons. (Other shipment special services may or may not be available per shipment for other reasons.)
	RestrictedShipmentSpecialServices []*SimpleType.ShipmentSpecialServiceType `xml:"RestrictedShipmentSpecialServices,omitempty"`

	FedExEuropeFirstOrigin bool `xml:"FedExEuropeFirstOrigin,omitempty"`
}
