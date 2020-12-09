package SimpleType

// Specifies the role that identifies the permissions the accessor of the pending shipment.
type AccessorRoleType string

const (
	AccessorRoleTypeSHIPMENT_COMPLETOR AccessorRoleType = "SHIPMENT_COMPLETOR"

	AccessorRoleTypeSHIPMENT_INITIATOR AccessorRoleType = "SHIPMENT_INITIATOR"
)
