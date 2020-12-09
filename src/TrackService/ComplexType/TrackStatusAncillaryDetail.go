package ComplexType

type TrackStatusAncillaryDetail struct {
	Reason string `xml:"Reason,omitempty"`

	ReasonDescription string `xml:"ReasonDescription,omitempty"`

	Action string `xml:"Action,omitempty"`

	ActionDescription string `xml:"ActionDescription,omitempty"`
}
