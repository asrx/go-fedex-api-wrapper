package SimpleType


// The description that FedEx uses for a given appointment window.
type AppointmentWindowType string

const (
	AppointmentWindowTypeAFTERNOON AppointmentWindowType = "AFTERNOON"

	AppointmentWindowTypeLATE_AFTERNOON AppointmentWindowType = "LATE_AFTERNOON"

	AppointmentWindowTypeMID_DAY AppointmentWindowType = "MID_DAY"

	AppointmentWindowTypeMORNING AppointmentWindowType = "MORNING"
)
