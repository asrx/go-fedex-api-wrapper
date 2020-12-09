package ComplexType

type ParsedAddressPartsDetail struct {
	ParsedStreetLine *ParsedStreetLineDetail `xml:"ParsedStreetLine,omitempty"`

	// The postal code specified in a form that is supported by USPS as base, secondary and tertiary.
	ParsedPostalCode *ParsedPostalCodeDetail `xml:"ParsedPostalCode,omitempty"`
}
