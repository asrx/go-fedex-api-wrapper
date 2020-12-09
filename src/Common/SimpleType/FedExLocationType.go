package SimpleType

// Identifies a kind of FedEx facility.
type FedExLocationType string

const (
	FedExLocationTypeFEDEX_EXPRESS_STATION FedExLocationType = "FEDEX_EXPRESS_STATION"

	FedExLocationTypeFEDEX_FACILITY FedExLocationType = "FEDEX_FACILITY"

	FedExLocationTypeFEDEX_FREIGHT_SERVICE_CENTER FedExLocationType = "FEDEX_FREIGHT_SERVICE_CENTER"

	FedExLocationTypeFEDEX_GROUND_TERMINAL FedExLocationType = "FEDEX_GROUND_TERMINAL"

	FedExLocationTypeFEDEX_HOME_DELIVERY_STATION FedExLocationType = "FEDEX_HOME_DELIVERY_STATION"

	FedExLocationTypeFEDEX_OFFICE FedExLocationType = "FEDEX_OFFICE"

	FedExLocationTypeFEDEX_SHIPSITE FedExLocationType = "FEDEX_SHIPSITE"

	FedExLocationTypeFEDEX_SMART_POST_HUB FedExLocationType = "FEDEX_SMART_POST_HUB"
)