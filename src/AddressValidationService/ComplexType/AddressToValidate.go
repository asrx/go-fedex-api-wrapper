package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type AddressToValidate struct {

	// A reference id provided by the client.
	ClientReferenceId string `xml:"ClientReferenceId,omitempty"`

	Contact *ComplexType.Contact `xml:"Contact,omitempty"`

	Address *ComplexType.Address `xml:"Address,omitempty"`
}
