package SimpleType

// Specifies the types of Estimated Duties and Taxes to be included in a rate quotation for an international shipment.
type EdtRequestType string

const (
	EdtRequestTypeALL EdtRequestType = "ALL"

	EdtRequestTypeNONE EdtRequestType = "NONE"
)