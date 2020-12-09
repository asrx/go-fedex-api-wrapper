package ComplexType

type TrackReconciliation struct {

	// An identifier for this type of status.
	Status string `xml:"Status,omitempty"`

	// A human-readable description of this status.
	Description string `xml:"Description,omitempty"`
}
