package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type TrackingDocumentDispositionDetail struct {
	DispositionType *SimpleType.TrackingDocumentDispositionType `xml:"DispositionType,omitempty"`

	EMailDetail *TrackingDocumentEmailDetail `xml:"EMailDetail,omitempty"`

	// Specifies the information used to fax the document.
	FaxDetails []*FaxDetail `xml:"FaxDetails,omitempty"`
}
