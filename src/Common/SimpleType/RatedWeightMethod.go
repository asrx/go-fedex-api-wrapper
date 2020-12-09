package SimpleType

type RatedWeightMethod string

const (
	RatedWeightMethodACTUAL RatedWeightMethod = "ACTUAL"

	RatedWeightMethodAVERAGE_PACKAGE_WEIGHT_MINIMUM RatedWeightMethod = "AVERAGE_PACKAGE_WEIGHT_MINIMUM"

	RatedWeightMethodBALLOON RatedWeightMethod = "BALLOON"

	RatedWeightMethodDEFAULT_WEIGHT_APPLIED RatedWeightMethod = "DEFAULT_WEIGHT_APPLIED"

	RatedWeightMethodDIM RatedWeightMethod = "DIM"

	RatedWeightMethodFREIGHT_MINIMUM RatedWeightMethod = "FREIGHT_MINIMUM"

	RatedWeightMethodMIXED RatedWeightMethod = "MIXED"

	RatedWeightMethodOVERSIZE RatedWeightMethod = "OVERSIZE"

	RatedWeightMethodOVERSIZE_1 RatedWeightMethod = "OVERSIZE_1"

	RatedWeightMethodOVERSIZE_2 RatedWeightMethod = "OVERSIZE_2"

	RatedWeightMethodOVERSIZE_3 RatedWeightMethod = "OVERSIZE_3"

	RatedWeightMethodPACKAGING_MINIMUM RatedWeightMethod = "PACKAGING_MINIMUM"

	RatedWeightMethodWEIGHT_BREAK RatedWeightMethod = "WEIGHT_BREAK"
)

