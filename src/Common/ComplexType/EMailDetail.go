package ComplexType

type EMailDetail struct {
	EmailAddress string `xml:"EmailAddress,omitempty"`

	// Specifies the name associated with the email address.
	Name string `xml:"Name,omitempty"`
}
