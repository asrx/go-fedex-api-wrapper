package SimpleType

// Specifies how to organize all shipping documents of the same type.
type ShippingDocumentGroupingType string

const (
	ShippingDocumentGroupingTypeCONSOLIDATED_BY_DOCUMENT_TYPE ShippingDocumentGroupingType = "CONSOLIDATED_BY_DOCUMENT_TYPE"

	ShippingDocumentGroupingTypeINDIVIDUAL ShippingDocumentGroupingType = "INDIVIDUAL"
)
