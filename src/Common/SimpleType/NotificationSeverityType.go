package SimpleType


type NotificationSeverityType string

const (
	NotificationSeverityTypeERROR NotificationSeverityType = "ERROR"

	NotificationSeverityTypeFAILURE NotificationSeverityType = "FAILURE"

	NotificationSeverityTypeNOTE NotificationSeverityType = "NOTE"

	NotificationSeverityTypeSUCCESS NotificationSeverityType = "SUCCESS"

	NotificationSeverityTypeWARNING NotificationSeverityType = "WARNING"
)
