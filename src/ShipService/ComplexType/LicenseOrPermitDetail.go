package ComplexType


type LicenseOrPermitDetail struct {

	// License or Permit Number.
	Number string `xml:"Number,omitempty"`

	// Specifies the effective date of the license.
	EffectiveDate string `xml:"EffectiveDate,omitempty"`

	// Specifies the expiration date of the license.
	ExpirationDate string `xml:"ExpirationDate,omitempty"`
}
