package ComplexType

type NotificationParameter struct {

	// Identifies the type of data contained in Value (e.g. SERVICE_TYPE, PACKAGE_SEQUENCE, etc..).
	Id string `xml:"Id,omitempty"`

	// The value of the parameter (e.g. PRIORITY_OVERNIGHT, 2, etc..).
	Value string `xml:"Value,omitempty"`
}
