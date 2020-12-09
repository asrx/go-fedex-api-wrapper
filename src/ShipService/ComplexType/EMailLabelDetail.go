package ComplexType

type EMailLabelDetail struct {

	// Content of the email message.
	Message string `xml:"Message,omitempty"`

	Recipients []*EMailRecipient `xml:"Recipients,omitempty"`
}
