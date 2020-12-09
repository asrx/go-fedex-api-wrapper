package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type PriorityAlertDetail struct {
	EnhancementTypes []*SimpleType.PriorityAlertEnhancementType `xml:"EnhancementTypes,omitempty"`

	Content []string `xml:"Content,omitempty"`
}
