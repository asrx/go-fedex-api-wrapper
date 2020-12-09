package SimpleType


type CustomsRoleType string

const (
	CustomsRoleTypeEXPORTER CustomsRoleType = "EXPORTER"

	CustomsRoleTypeIMPORTER CustomsRoleType = "IMPORTER"

	CustomsRoleTypeLEGAL_AGENT CustomsRoleType = "LEGAL_AGENT"

	CustomsRoleTypePRODUCER CustomsRoleType = "PRODUCER"
)
