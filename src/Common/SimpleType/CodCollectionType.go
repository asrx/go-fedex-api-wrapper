package SimpleType


// Identifies the type of funds FedEx should collect upon shipment delivery.
type CodCollectionType string

const (
	CodCollectionTypeANY CodCollectionType = "ANY"

	CodCollectionTypeCASH CodCollectionType = "CASH"

	CodCollectionTypeGUARANTEED_FUNDS CodCollectionType = "GUARANTEED_FUNDS"
)