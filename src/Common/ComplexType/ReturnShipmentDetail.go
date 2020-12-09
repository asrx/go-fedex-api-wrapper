package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ReturnShipmentDetail struct {
	ReturnType *SimpleType.ReturnType `xml:"ReturnType,omitempty"`

	Rma *Rma `xml:"Rma,omitempty"`

	ReturnEMailDetail *ReturnEMailDetail `xml:"ReturnEMailDetail,omitempty"`

	ReturnAssociation *ReturnAssociationDetail `xml:"ReturnAssociation,omitempty"`
}
