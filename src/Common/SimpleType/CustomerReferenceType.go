package SimpleType


type CustomerReferenceType string

const (
	CustomerReferenceTypeBILL_OF_LADING CustomerReferenceType = "BILL_OF_LADING"

	CustomerReferenceTypeCUSTOMER_REFERENCE CustomerReferenceType = "CUSTOMER_REFERENCE"

	CustomerReferenceTypeDEPARTMENT_NUMBER CustomerReferenceType = "DEPARTMENT_NUMBER"

	CustomerReferenceTypeELECTRONIC_PRODUCT_CODE CustomerReferenceType = "ELECTRONIC_PRODUCT_CODE"

	CustomerReferenceTypeINTRACOUNTRY_REGULATORY_REFERENCE CustomerReferenceType = "INTRACOUNTRY_REGULATORY_REFERENCE"

	CustomerReferenceTypeINVOICE_NUMBER CustomerReferenceType = "INVOICE_NUMBER"

	CustomerReferenceTypePACKING_SLIP_NUMBER CustomerReferenceType = "PACKING_SLIP_NUMBER"

	CustomerReferenceTypeP_O_NUMBER CustomerReferenceType = "P_O_NUMBER"

	CustomerReferenceTypeRMA_ASSOCIATION CustomerReferenceType = "RMA_ASSOCIATION"

	CustomerReferenceTypeSHIPMENT_INTEGRITY CustomerReferenceType = "SHIPMENT_INTEGRITY"

	CustomerReferenceTypeSTORE_NUMBER CustomerReferenceType = "STORE_NUMBER"
)