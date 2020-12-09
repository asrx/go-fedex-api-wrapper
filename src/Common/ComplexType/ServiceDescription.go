package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ServiceDescription struct {
	ServiceType *SimpleType.ServiceType `xml:"ServiceType,omitempty"`

	Code string `xml:"Code,omitempty"`

	// Branded, translated, and/or localized names for this service.
	Names []*ProductName `xml:"Names,omitempty"`

	Description string `xml:"Description,omitempty"`

	AstraDescription string `xml:"AstraDescription,omitempty"`
}
