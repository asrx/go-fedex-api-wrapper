package SimpleType


// Specifies the image format used for a shipping document.
type ShippingDocumentImageType string

const (
	ShippingDocumentImageTypeEPL2 ShippingDocumentImageType = "EPL2"

	ShippingDocumentImageTypePDF ShippingDocumentImageType = "PDF"

	ShippingDocumentImageTypePNG ShippingDocumentImageType = "PNG"

	ShippingDocumentImageTypeZPLII ShippingDocumentImageType = "ZPLII"
)
