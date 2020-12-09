package SimpleType

// Specifies how different the address returned is from the address provided. This difference can be because the address is cannonialised to match the address specification standard set by USPS.
type OperationalAddressStateType string

const (
	OperationalAddressStateTypeNORMALIZED OperationalAddressStateType = "NORMALIZED"

	OperationalAddressStateTypeRAW OperationalAddressStateType = "RAW"

	OperationalAddressStateTypeSTANDARDIZED OperationalAddressStateType = "STANDARDIZED"
)
