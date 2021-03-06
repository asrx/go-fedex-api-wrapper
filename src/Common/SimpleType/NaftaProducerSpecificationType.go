package SimpleType

type NaftaProducerSpecificationType string

const (
	NaftaProducerSpecificationTypeAVAILABLE_UPON_REQUEST NaftaProducerSpecificationType = "AVAILABLE_UPON_REQUEST"

	NaftaProducerSpecificationTypeMULTIPLE_SPECIFIED NaftaProducerSpecificationType = "MULTIPLE_SPECIFIED"

	NaftaProducerSpecificationTypeSAME NaftaProducerSpecificationType = "SAME"

	NaftaProducerSpecificationTypeSINGLE_SPECIFIED NaftaProducerSpecificationType = "SINGLE_SPECIFIED"

	NaftaProducerSpecificationTypeUNKNOWN NaftaProducerSpecificationType = "UNKNOWN"
)