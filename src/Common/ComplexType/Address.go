package ComplexType


type Address struct {

	// Combination of number, street name, etc. At least one line is required for a valid physical address; empty lines should not be included.
	StreetLines []string `xml:"StreetLines,omitempty"`

	// Name of city, town, etc.
	City string `xml:"City,omitempty"`

	// Identifying abbreviation for US state, Canada province, etc. Format and presence of this field will vary, depending on country.
	StateOrProvinceCode string `xml:"StateOrProvinceCode,omitempty"`

	// Identification of a region (usually small) for mail/package delivery. Format and presence of this field will vary, depending on country.
	PostalCode string `xml:"PostalCode,omitempty"`

	// Relevant only to addresses in Puerto Rico.
	UrbanizationCode string `xml:"UrbanizationCode,omitempty"`

	// The two-letter code used to identify a country.
	CountryCode string `xml:"CountryCode,omitempty"`

	// The fully spelt out name of a country.
	CountryName string `xml:"CountryName,omitempty"`

	// Indicates whether this address residential (as opposed to commercial).
	Residential bool `xml:"Residential,omitempty"`

	// The geographic coordinates cooresponding to this address.
	GeographicCoordinates string `xml:"GeographicCoordinates,omitempty"`
}
