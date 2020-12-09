package ComplexType

type ArtifactPart struct {

	// Specifies the sequence of this part of the artifact.
	SequenceNumber *uint `xml:"SequenceNumber,omitempty"`

	// Specifies the contents of this retrieved artifact part.
	Contents []byte `xml:"Contents,omitempty"`
}

