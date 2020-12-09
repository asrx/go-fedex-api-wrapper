package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type SignatureImageDetail struct {
	Image []byte `xml:"Image,omitempty"`

	Notifications []*ComplexType.Notification `xml:"Notifications,omitempty"`
}
