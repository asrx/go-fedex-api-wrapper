package SimpleType


// This identifies some of the document types recognized by Enterprise Document Management Service.
type EnterpriseDocumentType string

const (
	EnterpriseDocumentTypeAIR_WAYBILL EnterpriseDocumentType = "AIR_WAYBILL"

	EnterpriseDocumentTypeCERTIFICATE_OF_ORIGIN EnterpriseDocumentType = "CERTIFICATE_OF_ORIGIN"

	EnterpriseDocumentTypeCOMMERCIAL_INVOICE EnterpriseDocumentType = "COMMERCIAL_INVOICE"

	EnterpriseDocumentTypeNAFTA_CERTIFICATE_OF_ORIGIN EnterpriseDocumentType = "NAFTA_CERTIFICATE_OF_ORIGIN"

	EnterpriseDocumentTypePRO_FORMA_INVOICE EnterpriseDocumentType = "PRO_FORMA_INVOICE"
)
