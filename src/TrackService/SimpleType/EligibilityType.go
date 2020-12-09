package SimpleType


// Specifies different values of eligibility status
type EligibilityType string

const (
	EligibilityTypeELIGIBLE EligibilityType = "ELIGIBLE"

	EligibilityTypeINELIGIBLE EligibilityType = "INELIGIBLE"

	EligibilityTypePOSSIBLY_ELIGIBLE EligibilityType = "POSSIBLY_ELIGIBLE"
)
