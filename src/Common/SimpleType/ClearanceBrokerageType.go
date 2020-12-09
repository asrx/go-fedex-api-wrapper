package SimpleType


// Specifies the type of brokerage to be applied to a shipment.
type ClearanceBrokerageType string

const (
	ClearanceBrokerageTypeBROKER_INCLUSIVE ClearanceBrokerageType = "BROKER_INCLUSIVE"

	ClearanceBrokerageTypeBROKER_INCLUSIVE_NON_RESIDENT_IMPORTER ClearanceBrokerageType = "BROKER_INCLUSIVE_NON_RESIDENT_IMPORTER"

	ClearanceBrokerageTypeBROKER_SELECT ClearanceBrokerageType = "BROKER_SELECT"

	ClearanceBrokerageTypeBROKER_SELECT_NON_RESIDENT_IMPORTER ClearanceBrokerageType = "BROKER_SELECT_NON_RESIDENT_IMPORTER"

	ClearanceBrokerageTypeBROKER_UNASSIGNED ClearanceBrokerageType = "BROKER_UNASSIGNED"
)