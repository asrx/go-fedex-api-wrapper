package SimpleType


type EMailNotificationRecipientType string

const (
	EMailNotificationRecipientTypeBROKER EMailNotificationRecipientType = "BROKER"

	EMailNotificationRecipientTypeOTHER EMailNotificationRecipientType = "OTHER"

	EMailNotificationRecipientTypeRECIPIENT EMailNotificationRecipientType = "RECIPIENT"

	EMailNotificationRecipientTypeSHIPPER EMailNotificationRecipientType = "SHIPPER"
)