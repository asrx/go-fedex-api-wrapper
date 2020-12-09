package SimpleType

// Specifies the type of rate quote
type FreightRateQuoteType string

const (
	FreightRateQuoteTypeAUTOMATED FreightRateQuoteType = "AUTOMATED"

	FreightRateQuoteTypeMANUAL FreightRateQuoteType = "MANUAL"
)