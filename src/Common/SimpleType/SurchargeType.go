package SimpleType


type SurchargeType string

const (
	SurchargeTypeACCOUNT_NUMBER_PROCESSING_FEE SurchargeType = "ACCOUNT_NUMBER_PROCESSING_FEE"

	SurchargeTypeADDITIONAL_HANDLING SurchargeType = "ADDITIONAL_HANDLING"

	SurchargeTypeADDRESS_CORRECTION SurchargeType = "ADDRESS_CORRECTION"

	SurchargeTypeANCILLARY_FEE SurchargeType = "ANCILLARY_FEE"

	SurchargeTypeAPPOINTMENT_DELIVERY SurchargeType = "APPOINTMENT_DELIVERY"

	SurchargeTypeBROKER_SELECT_OPTION SurchargeType = "BROKER_SELECT_OPTION"

	SurchargeTypeCANADIAN_DESTINATION SurchargeType = "CANADIAN_DESTINATION"

	SurchargeTypeCLEARANCE_ENTRY_FEE SurchargeType = "CLEARANCE_ENTRY_FEE"

	SurchargeTypeCOD SurchargeType = "COD"

	SurchargeTypeCUT_FLOWERS SurchargeType = "CUT_FLOWERS"

	SurchargeTypeDANGEROUS_GOODS SurchargeType = "DANGEROUS_GOODS"

	SurchargeTypeDELIVERY_AREA SurchargeType = "DELIVERY_AREA"

	SurchargeTypeDELIVERY_CONFIRMATION SurchargeType = "DELIVERY_CONFIRMATION"

	SurchargeTypeDOCUMENTATION_FEE SurchargeType = "DOCUMENTATION_FEE"

	SurchargeTypeDRY_ICE SurchargeType = "DRY_ICE"

	SurchargeTypeEMAIL_LABEL SurchargeType = "EMAIL_LABEL"

	SurchargeTypeEUROPE_FIRST SurchargeType = "EUROPE_FIRST"

	SurchargeTypeEXCESS_VALUE SurchargeType = "EXCESS_VALUE"

	SurchargeTypeEXHIBITION SurchargeType = "EXHIBITION"

	SurchargeTypeEXPORT SurchargeType = "EXPORT"

	SurchargeTypeEXTRA_SURFACE_HANDLING_CHARGE SurchargeType = "EXTRA_SURFACE_HANDLING_CHARGE"

	SurchargeTypeEXTREME_LENGTH SurchargeType = "EXTREME_LENGTH"

	SurchargeTypeFEDEX_INTRACOUNTRY_FEES SurchargeType = "FEDEX_INTRACOUNTRY_FEES"

	SurchargeTypeFEDEX_TAG SurchargeType = "FEDEX_TAG"

	SurchargeTypeFICE SurchargeType = "FICE"

	SurchargeTypeFLATBED SurchargeType = "FLATBED"

	SurchargeTypeFREIGHT_GUARANTEE SurchargeType = "FREIGHT_GUARANTEE"

	SurchargeTypeFREIGHT_ON_VALUE SurchargeType = "FREIGHT_ON_VALUE"

	SurchargeTypeFREIGHT_TO_COLLECT SurchargeType = "FREIGHT_TO_COLLECT"

	SurchargeTypeFUEL SurchargeType = "FUEL"

	SurchargeTypeHOLD_AT_LOCATION SurchargeType = "HOLD_AT_LOCATION"

	SurchargeTypeHOME_DELIVERY_APPOINTMENT SurchargeType = "HOME_DELIVERY_APPOINTMENT"

	SurchargeTypeHOME_DELIVERY_DATE_CERTAIN SurchargeType = "HOME_DELIVERY_DATE_CERTAIN"

	SurchargeTypeHOME_DELIVERY_EVENING SurchargeType = "HOME_DELIVERY_EVENING"

	SurchargeTypeINSIDE_DELIVERY SurchargeType = "INSIDE_DELIVERY"

	SurchargeTypeINSIDE_PICKUP SurchargeType = "INSIDE_PICKUP"

	SurchargeTypeINSURED_VALUE SurchargeType = "INSURED_VALUE"

	SurchargeTypeINTERHAWAII SurchargeType = "INTERHAWAII"

	SurchargeTypeLIFTGATE_DELIVERY SurchargeType = "LIFTGATE_DELIVERY"

	SurchargeTypeLIFTGATE_PICKUP SurchargeType = "LIFTGATE_PICKUP"

	SurchargeTypeLIMITED_ACCESS_DELIVERY SurchargeType = "LIMITED_ACCESS_DELIVERY"

	SurchargeTypeLIMITED_ACCESS_PICKUP SurchargeType = "LIMITED_ACCESS_PICKUP"

	SurchargeTypeMETRO_DELIVERY SurchargeType = "METRO_DELIVERY"

	SurchargeTypeMETRO_PICKUP SurchargeType = "METRO_PICKUP"

	SurchargeTypeNON_MACHINABLE SurchargeType = "NON_MACHINABLE"

	SurchargeTypeOFFSHORE SurchargeType = "OFFSHORE"

	SurchargeTypeON_CALL_PICKUP SurchargeType = "ON_CALL_PICKUP"

	SurchargeTypeON_DEMAND_CARE SurchargeType = "ON_DEMAND_CARE"

	SurchargeTypeOTHER SurchargeType = "OTHER"

	SurchargeTypeOUT_OF_DELIVERY_AREA SurchargeType = "OUT_OF_DELIVERY_AREA"

	SurchargeTypeOUT_OF_PICKUP_AREA SurchargeType = "OUT_OF_PICKUP_AREA"

	SurchargeTypeOVERSIZE SurchargeType = "OVERSIZE"

	SurchargeTypeOVER_DIMENSION SurchargeType = "OVER_DIMENSION"

	SurchargeTypeOVER_LENGTH SurchargeType = "OVER_LENGTH"

	SurchargeTypePIECE_COUNT_VERIFICATION SurchargeType = "PIECE_COUNT_VERIFICATION"

	SurchargeTypePRE_DELIVERY_NOTIFICATION SurchargeType = "PRE_DELIVERY_NOTIFICATION"

	SurchargeTypePRIORITY_ALERT SurchargeType = "PRIORITY_ALERT"

	SurchargeTypePROTECTION_FROM_FREEZING SurchargeType = "PROTECTION_FROM_FREEZING"

	SurchargeTypeREGIONAL_MALL_DELIVERY SurchargeType = "REGIONAL_MALL_DELIVERY"

	SurchargeTypeREGIONAL_MALL_PICKUP SurchargeType = "REGIONAL_MALL_PICKUP"

	SurchargeTypeREROUTE SurchargeType = "REROUTE"

	SurchargeTypeRESCHEDULE SurchargeType = "RESCHEDULE"

	SurchargeTypeRESIDENTIAL_DELIVERY SurchargeType = "RESIDENTIAL_DELIVERY"

	SurchargeTypeRESIDENTIAL_PICKUP SurchargeType = "RESIDENTIAL_PICKUP"

	SurchargeTypeRETURN_LABEL SurchargeType = "RETURN_LABEL"

	SurchargeTypeSATURDAY_DELIVERY SurchargeType = "SATURDAY_DELIVERY"

	SurchargeTypeSATURDAY_PICKUP SurchargeType = "SATURDAY_PICKUP"

	SurchargeTypeSIGNATURE_OPTION SurchargeType = "SIGNATURE_OPTION"

	SurchargeTypeTARP SurchargeType = "TARP"

	SurchargeTypeTHIRD_PARTY_CONSIGNEE SurchargeType = "THIRD_PARTY_CONSIGNEE"

	SurchargeTypeTRANSMART_SERVICE_FEE SurchargeType = "TRANSMART_SERVICE_FEE"
)
