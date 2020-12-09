package SimpleType

// Specifies the address classification (business vs. residential)
type FedExAddressClassificationType string

const (
	FedExAddressClassificationTypeBUSINESS FedExAddressClassificationType = "BUSINESS"

	FedExAddressClassificationTypeMIXED FedExAddressClassificationType = "MIXED"

	FedExAddressClassificationTypeRESIDENTIAL FedExAddressClassificationType = "RESIDENTIAL"

	FedExAddressClassificationTypeUNKNOWN FedExAddressClassificationType = "UNKNOWN"
)
