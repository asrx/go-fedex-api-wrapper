package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type SignatureOptionDetail struct {
	OptionType *SimpleType.SignatureOptionType `xml:"OptionType,omitempty"`

	SignatureReleaseNumber string `xml:"SignatureReleaseNumber,omitempty"`
}
