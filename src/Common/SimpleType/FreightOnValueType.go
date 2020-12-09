package SimpleType

// Identifies responsibilities with respect to loss, damage, etc.
type FreightOnValueType string

const (
	FreightOnValueTypeCARRIER_RISK FreightOnValueType = "CARRIER_RISK"

	FreightOnValueTypeOWN_RISK FreightOnValueType = "OWN_RISK"
)