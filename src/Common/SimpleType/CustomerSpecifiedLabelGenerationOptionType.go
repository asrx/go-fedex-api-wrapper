package SimpleType


type CustomerSpecifiedLabelGenerationOptionType string

const (
	CustomerSpecifiedLabelGenerationOptionTypeCONTENT_ON_SHIPPING_LABEL_ONLY CustomerSpecifiedLabelGenerationOptionType = "CONTENT_ON_SHIPPING_LABEL_ONLY"

	CustomerSpecifiedLabelGenerationOptionTypeCONTENT_ON_SHIPPING_LABEL_PREFERRED CustomerSpecifiedLabelGenerationOptionType = "CONTENT_ON_SHIPPING_LABEL_PREFERRED"

	CustomerSpecifiedLabelGenerationOptionTypeCONTENT_ON_SUPPLEMENTAL_LABEL_ONLY CustomerSpecifiedLabelGenerationOptionType = "CONTENT_ON_SUPPLEMENTAL_LABEL_ONLY"
)