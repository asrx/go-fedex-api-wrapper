package SimpleType

type PaymentType string

const (
	PaymentTypeSENDER PaymentType = "SENDER"
	PaymentTypeTHIRD_PARTY PaymentType = "THIRD_PARTY"
)