package SimpleType

// Selects the value from a set of rate data to which the percentage is applied.
type RateElementBasisType string

const (
	RateElementBasisTypeBASE_CHARGE RateElementBasisType = "BASE_CHARGE"

	RateElementBasisTypeNET_CHARGE RateElementBasisType = "NET_CHARGE"

	RateElementBasisTypeNET_CHARGE_EXCLUDING_TAXES RateElementBasisType = "NET_CHARGE_EXCLUDING_TAXES"

	RateElementBasisTypeNET_FREIGHT RateElementBasisType = "NET_FREIGHT"
)
