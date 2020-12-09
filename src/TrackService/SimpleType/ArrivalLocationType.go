package SimpleType


// Identifies where a tracking event occurs.
type ArrivalLocationType string

const (
	ArrivalLocationTypeAIRPORT ArrivalLocationType = "AIRPORT"

	ArrivalLocationTypeCUSTOMER ArrivalLocationType = "CUSTOMER"

	ArrivalLocationTypeCUSTOMS_BROKER ArrivalLocationType = "CUSTOMS_BROKER"

	ArrivalLocationTypeDELIVERY_LOCATION ArrivalLocationType = "DELIVERY_LOCATION"

	ArrivalLocationTypeDESTINATION_AIRPORT ArrivalLocationType = "DESTINATION_AIRPORT"

	ArrivalLocationTypeDESTINATION_FEDEX_FACILITY ArrivalLocationType = "DESTINATION_FEDEX_FACILITY"

	ArrivalLocationTypeDROP_BOX ArrivalLocationType = "DROP_BOX"

	ArrivalLocationTypeENROUTE ArrivalLocationType = "ENROUTE"

	ArrivalLocationTypeFEDEX_FACILITY ArrivalLocationType = "FEDEX_FACILITY"

	ArrivalLocationTypeFEDEX_OFFICE_LOCATION ArrivalLocationType = "FEDEX_OFFICE_LOCATION"

	ArrivalLocationTypeINTERLINE_CARRIER ArrivalLocationType = "INTERLINE_CARRIER"

	ArrivalLocationTypeNON_FEDEX_FACILITY ArrivalLocationType = "NON_FEDEX_FACILITY"

	ArrivalLocationTypeORIGIN_AIRPORT ArrivalLocationType = "ORIGIN_AIRPORT"

	ArrivalLocationTypeORIGIN_FEDEX_FACILITY ArrivalLocationType = "ORIGIN_FEDEX_FACILITY"

	ArrivalLocationTypePICKUP_LOCATION ArrivalLocationType = "PICKUP_LOCATION"

	ArrivalLocationTypePLANE ArrivalLocationType = "PLANE"

	ArrivalLocationTypePORT_OF_ENTRY ArrivalLocationType = "PORT_OF_ENTRY"

	ArrivalLocationTypeSHIP_AND_GET_LOCATION ArrivalLocationType = "SHIP_AND_GET_LOCATION"

	ArrivalLocationTypeSORT_FACILITY ArrivalLocationType = "SORT_FACILITY"

	ArrivalLocationTypeTURNPOINT ArrivalLocationType = "TURNPOINT"

	ArrivalLocationTypeVEHICLE ArrivalLocationType = "VEHICLE"
)
