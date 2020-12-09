package ComplexType

type Localization struct {

	// Two-letter code for language (e.g. EN, FR, etc.)
	LanguageCode string `xml:"LanguageCode,omitempty"`

	// Two-letter code for the region (e.g. us, ca, etc..).
	LocaleCode string `xml:"LocaleCode,omitempty"`
}
