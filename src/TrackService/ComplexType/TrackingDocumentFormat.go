package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackingDocumentFormat struct {
	Dispositions []*TrackingDocumentDispositionDetail `xml:"Dispositions,omitempty"`

	Grouping *SimpleType.TrackingDocumentGroupingType `xml:"Grouping,omitempty"`

	ImageType *SimpleType.TrackingDocumentImageType `xml:"ImageType,omitempty"`

	// The localization for the generated document.
	Localization *ComplexType.Localization `xml:"Localization,omitempty"`
}
