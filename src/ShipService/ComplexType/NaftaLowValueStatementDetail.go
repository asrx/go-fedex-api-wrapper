package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type NaftaLowValueStatementDetail struct {

	// Specifies the NAFTA statement role.
	Role *SimpleType.CustomsRoleType `xml:"Role,omitempty"`
}
