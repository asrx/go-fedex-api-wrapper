package ComplexType


type CustomerExceptionRequestDetail struct {

	// Unique identifier for the customer exception request.
	Id string `xml:"Id,omitempty"`

	StatusCode string `xml:"StatusCode,omitempty"`

	StatusDescription string `xml:"StatusDescription,omitempty"`

	CreateTime string `xml:"CreateTime,omitempty"`
}
