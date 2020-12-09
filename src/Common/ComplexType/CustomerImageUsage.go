package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type CustomerImageUsage struct {
	Type *SimpleType.CustomerImageUsageType `xml:"Type,omitempty"`

	Id *SimpleType.ImageId `xml:"Id,omitempty"`
}
