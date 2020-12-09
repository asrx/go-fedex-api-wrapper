package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type CustomsDeclarationStatementDetail struct {

	// This indicates the different statements, declarations, acts, and/or certifications that apply to this shipment.
	Types []*SimpleType.CustomsDeclarationStatementType `xml:"Types,omitempty"`

	// Specifies the NAFTA low value statement information.
	NaftaLowValueStatementDetail *NaftaLowValueStatementDetail `xml:"NaftaLowValueStatementDetail,omitempty"`
}
