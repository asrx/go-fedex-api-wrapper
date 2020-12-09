package SimpleType

// Specifies the type of license that the recipient of the alcohol shipment has.
type AlcoholRecipientType string

const (
	AlcoholRecipientTypeCONSUMER AlcoholRecipientType = "CONSUMER"

	AlcoholRecipientTypeLICENSEE AlcoholRecipientType = "LICENSEE"
)
