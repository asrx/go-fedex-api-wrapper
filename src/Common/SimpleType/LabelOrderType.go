package SimpleType


// Specifies the order in which the labels will be returned
type LabelOrderType string

const (
	LabelOrderTypeSHIPPING_LABEL_FIRST LabelOrderType = "SHIPPING_LABEL_FIRST"

	LabelOrderTypeSHIPPING_LABEL_LAST LabelOrderType = "SHIPPING_LABEL_LAST"
)