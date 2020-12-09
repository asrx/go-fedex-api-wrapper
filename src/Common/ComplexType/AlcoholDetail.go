package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type AlcoholDetail struct {

	// The license type that the recipient of the alcohol package.
	RecipientType *SimpleType.AlcoholRecipientType `xml:"RecipientType,omitempty"`
}
