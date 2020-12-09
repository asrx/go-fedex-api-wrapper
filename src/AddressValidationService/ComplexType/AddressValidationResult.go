package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/AddressValidationService/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
)

type AddressValidationResult struct {

	// The client reference id for the validated address.
	ClientReferenceId string `xml:"ClientReferenceId,omitempty"`

	// Specifies the degree to SHARE service was able to cannonicalise the address provided, as per USPS standards and match it to an address already in the internal FedEx address repository.
	State *SimpleType.OperationalAddressStateType `xml:"State,omitempty"`

	Classification *SimpleType.FedExAddressClassificationType `xml:"Classification,omitempty"`

	EffectiveContact *ComplexType.Contact `xml:"EffectiveContact,omitempty"`

	EffectiveAddress *ComplexType.Address `xml:"EffectiveAddress,omitempty"`

	ParsedAddressPartsDetail *ParsedAddressPartsDetail `xml:"ParsedAddressPartsDetail,omitempty"`

	// Additional attributes about the validated address from FedEx address respository..
	Attributes []*AddressAttribute `xml:"Attributes,omitempty"`
}
