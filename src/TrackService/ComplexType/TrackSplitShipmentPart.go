package ComplexType


type TrackSplitShipmentPart struct {

	// The number of pieces in this part.
	PieceCount *uint `xml:"PieceCount,omitempty"`

	// The date and time this status began.
	Timestamp string `xml:"Timestamp,omitempty"`

	// A code that identifies this type of status.
	StatusCode string `xml:"StatusCode,omitempty"`

	// A human-readable description of this status.
	StatusDescription string `xml:"StatusDescription,omitempty"`
}
