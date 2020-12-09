package ComplexType

type WebAuthenticationCredential struct {

	// Identifying part of authentication credential. This value is provided by FedEx after registration
	Key string `xml:"Key,omitempty"`

	// Secret part of authentication key. This value is provided by FedEx after registration.
	Password string `xml:"Password,omitempty"`
}

