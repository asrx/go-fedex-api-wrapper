package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type DocTabContent struct {
	DocTabContentType *SimpleType.DocTabContentType `xml:"DocTabContentType,omitempty"`

	Zone001 *DocTabContentZone001 `xml:"Zone001,omitempty"`

	Barcoded *DocTabContentBarcoded `xml:"Barcoded,omitempty"`
}
