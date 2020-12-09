package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type Payment struct {
	PaymentType *SimpleType.PaymentType `xml:"PaymentType,omitempty"`

	Payor *Payor `xml:"Payor,omitempty"`
}
