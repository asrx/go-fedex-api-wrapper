package SimpleType

type TrackPossessionStatusType string

const (
	TrackPossessionStatusTypeBROKER TrackPossessionStatusType = "BROKER"

	TrackPossessionStatusTypeCARRIER TrackPossessionStatusType = "CARRIER"

	TrackPossessionStatusTypeCUSTOMS TrackPossessionStatusType = "CUSTOMS"

	TrackPossessionStatusTypeRECIPIENT TrackPossessionStatusType = "RECIPIENT"

	TrackPossessionStatusTypeSHIPPER TrackPossessionStatusType = "SHIPPER"

	TrackPossessionStatusTypeSPLIT_STATUS TrackPossessionStatusType = "SPLIT_STATUS"

	TrackPossessionStatusTypeTRANSFER_PARTNER TrackPossessionStatusType = "TRANSFER_PARTNER"
)
