package ComplexType


type DocumentPart struct {

	// The one-origin position of this part within a document.
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	// Graphic or printer commands for this image within a document.
	Content []byte `xml:"Content,omitempty"`
}
