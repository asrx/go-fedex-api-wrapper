package SimpleType


// The point where the delay is occurring ( e.g. Origin, Destination, Broker location).
type DelayPointType string

const (
	DelayPointTypeBROKER DelayPointType = "BROKER"

	DelayPointTypeDESTINATION DelayPointType = "DESTINATION"

	DelayPointTypeORIGIN DelayPointType = "ORIGIN"

	DelayPointTypeORIGIN_DESTINATION_PAIR DelayPointType = "ORIGIN_DESTINATION_PAIR"

	DelayPointTypePROOF_OF_DELIVERY_POINT DelayPointType = "PROOF_OF_DELIVERY_POINT"
)
