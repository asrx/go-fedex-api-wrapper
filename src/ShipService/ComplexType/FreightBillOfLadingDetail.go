package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

type FreightBillOfLadingDetail struct {
	Format *ComplexType.ShippingDocumentFormat `xml:"Format,omitempty"`
}
