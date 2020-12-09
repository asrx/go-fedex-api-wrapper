package ComplexType

type WebAuthenticationDetail struct {

	// This was renamed from cspCredential.
	ParentCredential *WebAuthenticationCredential `xml:"ParentCredential,omitempty"`

	// Credential used to authenticate a specific software application. This value is provided by FedEx after registration.
	UserCredential *WebAuthenticationCredential `xml:"UserCredential,omitempty"`
}

