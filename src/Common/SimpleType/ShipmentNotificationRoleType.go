package SimpleType

type ShipmentNotificationRoleType string

const (
	ShipmentNotificationRoleTypeBROKER ShipmentNotificationRoleType = "BROKER"

	ShipmentNotificationRoleTypeOTHER ShipmentNotificationRoleType = "OTHER"

	ShipmentNotificationRoleTypeRECIPIENT ShipmentNotificationRoleType = "RECIPIENT"

	ShipmentNotificationRoleTypeSHIPPER ShipmentNotificationRoleType = "SHIPPER"

	ShipmentNotificationRoleTypeTHIRD_PARTY ShipmentNotificationRoleType = "THIRD_PARTY"
)
