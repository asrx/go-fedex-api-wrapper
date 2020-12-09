package SimpleType


// These values are mutually exclusive; at most one of them can be attached to a SmartPost shipment.
type SmartPostAncillaryEndorsementType string

const (
	SmartPostAncillaryEndorsementTypeADDRESS_CORRECTION SmartPostAncillaryEndorsementType = "ADDRESS_CORRECTION"

	SmartPostAncillaryEndorsementTypeCARRIER_LEAVE_IF_NO_RESPONSE SmartPostAncillaryEndorsementType = "CARRIER_LEAVE_IF_NO_RESPONSE"

	SmartPostAncillaryEndorsementTypeCHANGE_SERVICE SmartPostAncillaryEndorsementType = "CHANGE_SERVICE"

	SmartPostAncillaryEndorsementTypeFORWARDING_SERVICE SmartPostAncillaryEndorsementType = "FORWARDING_SERVICE"

	SmartPostAncillaryEndorsementTypeRETURN_SERVICE SmartPostAncillaryEndorsementType = "RETURN_SERVICE"
)