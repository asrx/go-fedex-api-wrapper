package SimpleType


// Specifies how to return a shipping document to the caller.
type ShippingDocumentDispositionType string

const (
	ShippingDocumentDispositionTypeCONFIRMED ShippingDocumentDispositionType = "CONFIRMED"

	ShippingDocumentDispositionTypeDEFERRED_QUEUED ShippingDocumentDispositionType = "DEFERRED_QUEUED"

	ShippingDocumentDispositionTypeDEFERRED_RETURNED ShippingDocumentDispositionType = "DEFERRED_RETURNED"

	ShippingDocumentDispositionTypeDEFERRED_STORED ShippingDocumentDispositionType = "DEFERRED_STORED"

	ShippingDocumentDispositionTypeEMAILED ShippingDocumentDispositionType = "EMAILED"

	ShippingDocumentDispositionTypeQUEUED ShippingDocumentDispositionType = "QUEUED"

	ShippingDocumentDispositionTypeRETURNED ShippingDocumentDispositionType = "RETURNED"

	ShippingDocumentDispositionTypeSTORED ShippingDocumentDispositionType = "STORED"
)

