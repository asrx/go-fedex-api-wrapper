package SimpleType

// Indicates the role of the party submitting the transaction.
type FreightShipmentRoleType string

const (
	FreightShipmentRoleTypeCONSIGNEE FreightShipmentRoleType = "CONSIGNEE"

	FreightShipmentRoleTypeSHIPPER FreightShipmentRoleType = "SHIPPER"
)