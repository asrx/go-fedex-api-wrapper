package ComplexType


type TrackInformationNoteDetail struct {

	// A code that designates the type of informational message being returned.
	Code string `xml:"Code,omitempty"`

	// The informational message in human readable form.
	Description string `xml:"Description,omitempty"`
}
