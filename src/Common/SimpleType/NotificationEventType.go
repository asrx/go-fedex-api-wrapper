package SimpleType

type NotificationEventType string

const (
	NotificationEventTypeON_DELIVERY NotificationEventType = "ON_DELIVERY"

	NotificationEventTypeON_ESTIMATED_DELIVERY NotificationEventType = "ON_ESTIMATED_DELIVERY"

	NotificationEventTypeON_EXCEPTION NotificationEventType = "ON_EXCEPTION"

	NotificationEventTypeON_SHIPMENT NotificationEventType = "ON_SHIPMENT"

	NotificationEventTypeON_TENDER NotificationEventType = "ON_TENDER"
)