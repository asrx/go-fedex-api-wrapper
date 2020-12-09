package SimpleType

// Specifies the type of service scheduling offered from a Freight or National Freight Service Center to a customer-supplied address.
type FreightServiceSchedulingType string

const (
	FreightServiceSchedulingTypeLIMITED FreightServiceSchedulingType = "LIMITED"

	FreightServiceSchedulingTypeSTANDARD FreightServiceSchedulingType = "STANDARD"

	FreightServiceSchedulingTypeWILL_CALL FreightServiceSchedulingType = "WILL_CALL"
)