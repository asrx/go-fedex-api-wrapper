package SimpleType


// The type of delay this shipment will encounter.
type CommitmentDelayType string

const (
	CommitmentDelayTypeHOLIDAY CommitmentDelayType = "HOLIDAY"

	CommitmentDelayTypeNON_WORKDAY CommitmentDelayType = "NON_WORKDAY"

	CommitmentDelayTypeNO_CITY_DELIVERY CommitmentDelayType = "NO_CITY_DELIVERY"

	CommitmentDelayTypeNO_HOLD_AT_LOCATION CommitmentDelayType = "NO_HOLD_AT_LOCATION"

	CommitmentDelayTypeNO_LOCATION_DELIVERY CommitmentDelayType = "NO_LOCATION_DELIVERY"

	CommitmentDelayTypeNO_SERVICE_AREA_DELIVERY CommitmentDelayType = "NO_SERVICE_AREA_DELIVERY"

	CommitmentDelayTypeNO_SERVICE_AREA_SPECIAL_SERVICE_DELIVERY CommitmentDelayType = "NO_SERVICE_AREA_SPECIAL_SERVICE_DELIVERY"

	CommitmentDelayTypeNO_SPECIAL_SERVICE_DELIVERY CommitmentDelayType = "NO_SPECIAL_SERVICE_DELIVERY"

	CommitmentDelayTypeNO_ZIP_DELIVERY CommitmentDelayType = "NO_ZIP_DELIVERY"

	CommitmentDelayTypeWEEKEND CommitmentDelayType = "WEEKEND"

	CommitmentDelayTypeWEEKEND_SPECIAL CommitmentDelayType = "WEEKEND_SPECIAL"
)