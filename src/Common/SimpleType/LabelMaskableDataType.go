package SimpleType

// Names for data elements / areas which may be suppressed from printing on labels.
type LabelMaskableDataType string

const (
	LabelMaskableDataTypeCUSTOMS_VALUE LabelMaskableDataType = "CUSTOMS_VALUE"

	LabelMaskableDataTypeDIMENSIONS LabelMaskableDataType = "DIMENSIONS"

	LabelMaskableDataTypeDUTIES_AND_TAXES_PAYOR_ACCOUNT_NUMBER LabelMaskableDataType = "DUTIES_AND_TAXES_PAYOR_ACCOUNT_NUMBER"

	LabelMaskableDataTypeFREIGHT_PAYOR_ACCOUNT_NUMBER LabelMaskableDataType = "FREIGHT_PAYOR_ACCOUNT_NUMBER"

	LabelMaskableDataTypePACKAGE_SEQUENCE_AND_COUNT LabelMaskableDataType = "PACKAGE_SEQUENCE_AND_COUNT"

	LabelMaskableDataTypeSECONDARY_BARCODE LabelMaskableDataType = "SECONDARY_BARCODE"

	LabelMaskableDataTypeSHIPPER_ACCOUNT_NUMBER LabelMaskableDataType = "SHIPPER_ACCOUNT_NUMBER"

	LabelMaskableDataTypeSUPPLEMENTAL_LABEL_DOC_TAB LabelMaskableDataType = "SUPPLEMENTAL_LABEL_DOC_TAB"

	LabelMaskableDataTypeTERMS_AND_CONDITIONS LabelMaskableDataType = "TERMS_AND_CONDITIONS"

	LabelMaskableDataTypeTOTAL_WEIGHT LabelMaskableDataType = "TOTAL_WEIGHT"

	LabelMaskableDataTypeTRANSPORTATION_CHARGES_PAYOR_ACCOUNT_NUMBER LabelMaskableDataType = "TRANSPORTATION_CHARGES_PAYOR_ACCOUNT_NUMBER"
)
