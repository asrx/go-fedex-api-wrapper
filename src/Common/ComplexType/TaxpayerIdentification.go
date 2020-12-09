package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type TaxpayerIdentification struct {
	TinType *SimpleType.TinType `xml:"TinType,omitempty"`

	Number string `xml:"Number,omitempty"`

	// Identifies the usage of Tax Identification Number in Shipment processing
	Usage string `xml:"Usage,omitempty"`

	EffectiveDate string `xml:"EffectiveDate,omitempty"`

	ExpirationDate string `xml:"ExpirationDate,omitempty"`
}
