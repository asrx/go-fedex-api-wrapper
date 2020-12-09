package SimpleType


// The attribute of the shipment that caused the delay(e.g. Country, City, LocationId, Zip, service area, special handling )
type DelayLevelType string

const (
	DelayLevelTypeCITY DelayLevelType = "CITY"

	DelayLevelTypeCOUNTRY DelayLevelType = "COUNTRY"

	DelayLevelTypeLOCATION DelayLevelType = "LOCATION"

	DelayLevelTypePOSTAL_CODE DelayLevelType = "POSTAL_CODE"

	DelayLevelTypeSERVICE_AREA DelayLevelType = "SERVICE_AREA"

	DelayLevelTypeSERVICE_AREA_SPECIAL_SERVICE DelayLevelType = "SERVICE_AREA_SPECIAL_SERVICE"

	DelayLevelTypeSPECIAL_SERVICE DelayLevelType = "SPECIAL_SERVICE"
)
