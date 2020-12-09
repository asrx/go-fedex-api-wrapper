package ComplexType


type CurrencyExchangeRate struct {

	// The currency code for the original (converted FROM) currency.
	FromCurrency string `xml:"FromCurrency,omitempty"`

	// The currency code for the final (converted INTO) currency.
	IntoCurrency string `xml:"IntoCurrency,omitempty"`

	// Multiplier used to convert fromCurrency units to intoCurrency units.
	Rate float64 `xml:"Rate,omitempty"`
}