package ComplexType

type Contact struct {

	// Client provided identifier corresponding to this contact information.
	ContactId string `xml:"ContactId,omitempty"`

	// Identifies the contact person's name.
	PersonName string `xml:"PersonName,omitempty"`

	// Identifies the contact person's title.
	Title string `xml:"Title,omitempty"`

	// Identifies the company this contact is associated with.
	CompanyName string `xml:"CompanyName,omitempty"`

	// Identifies the phone number associated with this contact.
	PhoneNumber string `xml:"PhoneNumber,omitempty"`

	// Identifies the phone extension associated with this contact.
	PhoneExtension string `xml:"PhoneExtension,omitempty"`

	// Identifies a toll free number, if any, associated with this contact.
	TollFreePhoneNumber string `xml:"TollFreePhoneNumber,omitempty"`

	// Identifies the pager number associated with this contact.
	PagerNumber string `xml:"PagerNumber,omitempty"`

	// Identifies the fax number associated with this contact.
	FaxNumber string `xml:"FaxNumber,omitempty"`

	// Identifies the email address associated with this contact.
	EMailAddress string `xml:"EMailAddress,omitempty"`
}
