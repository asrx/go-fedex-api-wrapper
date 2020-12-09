package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DestinationControlDetail struct {
	StatementTypes []*SimpleType.DestinationControlStatementType `xml:"StatementTypes,omitempty"`

	// Comma-separated list of up to four country codes, required for DEPARTMENT_OF_STATE statement.
	DestinationCountries string `xml:"DestinationCountries,omitempty"`

	// Name of end user, required for DEPARTMENT_OF_STATE statement.
	EndUser string `xml:"EndUser,omitempty"`
}
