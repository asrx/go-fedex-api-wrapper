package SimpleType

type TrackPaymentType string

const (
	TrackPaymentTypeCASH_OR_CHECK_AT_DESTINATION TrackPaymentType = "CASH_OR_CHECK_AT_DESTINATION"

	TrackPaymentTypeCASH_OR_CHECK_AT_ORIGIN TrackPaymentType = "CASH_OR_CHECK_AT_ORIGIN"

	TrackPaymentTypeCREDIT_CARD_AT_DESTINATION TrackPaymentType = "CREDIT_CARD_AT_DESTINATION"

	TrackPaymentTypeCREDIT_CARD_AT_ORIGIN TrackPaymentType = "CREDIT_CARD_AT_ORIGIN"

	TrackPaymentTypeOTHER TrackPaymentType = "OTHER"

	TrackPaymentTypeRECIPIENT_ACCOUNT TrackPaymentType = "RECIPIENT_ACCOUNT"

	TrackPaymentTypeSHIPPER_ACCOUNT TrackPaymentType = "SHIPPER_ACCOUNT"

	TrackPaymentTypeTHIRD_PARTY_ACCOUNT TrackPaymentType = "THIRD_PARTY_ACCOUNT"
)

