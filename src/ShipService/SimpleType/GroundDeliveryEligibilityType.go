package SimpleType

// Represents features of FedEx Ground delivery for which the shipment is eligible.
type GroundDeliveryEligibilityType string

const (
	GroundDeliveryEligibilityTypeALTERNATE_DAY_SERVICE GroundDeliveryEligibilityType = "ALTERNATE_DAY_SERVICE"

	GroundDeliveryEligibilityTypeCARTAGE_AGENT_DELIVERY GroundDeliveryEligibilityType = "CARTAGE_AGENT_DELIVERY"

	GroundDeliveryEligibilityTypeSATURDAY_DELIVERY GroundDeliveryEligibilityType = "SATURDAY_DELIVERY"

	GroundDeliveryEligibilityTypeUSPS_DELIVERY GroundDeliveryEligibilityType = "USPS_DELIVERY"
)
