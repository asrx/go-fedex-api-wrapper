package ComplexType

type ShippingDocumentPart struct {

	// The one-origin position of this part within a document.
	DocumentPartSequenceNumber *uint `xml:"DocumentPartSequenceNumber,omitempty"`

	// Graphic or printer commands for this image within a document.
	Image []byte `xml:"Image,omitempty"`
}
