package SimpleType

// Identifies a particular reference identifier printed on a Freight bill of lading.
type PrintedReferenceType string

const (
	PrintedReferenceTypeBILL_OF_LADING PrintedReferenceType = "BILL_OF_LADING"

	PrintedReferenceTypeCONSIGNEE_ID_NUMBER PrintedReferenceType = "CONSIGNEE_ID_NUMBER"

	PrintedReferenceTypeSHIPPER_ID_NUMBER PrintedReferenceType = "SHIPPER_ID_NUMBER"
)
