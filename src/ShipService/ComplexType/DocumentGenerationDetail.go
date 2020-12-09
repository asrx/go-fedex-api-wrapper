package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type DocumentGenerationDetail struct {
	Type *SimpleType.EnterpriseDocumentType `xml:"Type,omitempty"`

	MinimumCopiesRequired *uint `xml:"MinimumCopiesRequired,omitempty"`

	Letterhead *SimpleType.RequirementType `xml:"Letterhead,omitempty"`

	ElectronicSignature *SimpleType.RequirementType `xml:"ElectronicSignature,omitempty"`
}
