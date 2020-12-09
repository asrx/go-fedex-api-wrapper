package ComplexType


type CleansedAddressAndLocationDetail struct {

	// This represents the internal FedEx-system recognized country code.
	CountryCode string `xml:"CountryCode,omitempty"`

	// This represents the internal FedEx-system recognized state or province code.
	StateOrProvinceCode string `xml:"StateOrProvinceCode,omitempty"`

	// This represents the internal FedEx-system recognized postal code.
	PostalCode string `xml:"PostalCode,omitempty"`

	ServiceArea string `xml:"ServiceArea,omitempty"`

	// The unique location identifier
	LocationId string `xml:"LocationId,omitempty"`

	// The op-co specific numeric identifier for a location
	LocationNumber int32 `xml:"LocationNumber,omitempty"`

	AirportId string `xml:"AirportId,omitempty"`
}
