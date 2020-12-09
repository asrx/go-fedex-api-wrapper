package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type BrokerDetail struct {
	Type *SimpleType.BrokerType `xml:"Type,omitempty"`

	Broker *Party `xml:"Broker,omitempty"`
}
