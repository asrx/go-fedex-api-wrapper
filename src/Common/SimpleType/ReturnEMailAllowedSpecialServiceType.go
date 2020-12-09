package SimpleType


// These values are used to control the availability of certain special services at the time when a customer uses the e-mail label link to create a return shipment.
type ReturnEMailAllowedSpecialServiceType string

const (
	ReturnEMailAllowedSpecialServiceTypeSATURDAY_DELIVERY ReturnEMailAllowedSpecialServiceType = "SATURDAY_DELIVERY"

	ReturnEMailAllowedSpecialServiceTypeSATURDAY_PICKUP ReturnEMailAllowedSpecialServiceType = "SATURDAY_PICKUP"
)