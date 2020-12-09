package ComplexType

type EdtCommodityTax struct {
	HarmonizedCode string `xml:"HarmonizedCode,omitempty"`

	Taxes []*EdtTaxDetail `xml:"Taxes,omitempty"`
}
