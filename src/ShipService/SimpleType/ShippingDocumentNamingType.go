package SimpleType

// Identifies the convention by which file names are constructed for STORED or DEFERRED documents.
type ShippingDocumentNamingType string

const (
	ShippingDocumentNamingTypeFAST ShippingDocumentNamingType = "FAST"

	ShippingDocumentNamingTypeLEGACY_FXRS ShippingDocumentNamingType = "LEGACY_FXRS"
)
