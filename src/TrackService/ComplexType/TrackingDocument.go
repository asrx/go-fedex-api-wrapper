package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
)

type TrackingDocument struct {
	Type *SimpleType.TrackingDocumentType `xml:"Type,omitempty"`

	Localizations []*ComplexType.Localization `xml:"Localizations,omitempty"`

	ImageType *SimpleType.TrackingDocumentImageType `xml:"ImageType,omitempty"`

	Resolution *uint `xml:"Resolution,omitempty"`

	Parts []*DocumentPart `xml:"Parts,omitempty"`
}
