package ComplexType


type Party struct {
	AccountNumber string `xml:"AccountNumber,omitempty"`

	Tins []*TaxpayerIdentification `xml:"Tins,omitempty"`

	Contact *Contact `xml:"Contact,omitempty"`

	Address *Address `xml:"Address,omitempty"`
}
