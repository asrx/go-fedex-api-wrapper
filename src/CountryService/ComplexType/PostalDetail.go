package ComplexType

type PostalDetail struct {
	CountryCode string `xml:"CountryCode,omitempty"`

	StateOrProvinceCode string `xml:"StateOrProvinceCode,omitempty"`

	CityFirstInitials string `xml:"CityFirstInitials,omitempty"`

	CleanedPostalCode string `xml:"CleanedPostalCode,omitempty"`

	LocationDescriptions []*LocationDescription `xml:"LocationDescriptions,omitempty"`
}
