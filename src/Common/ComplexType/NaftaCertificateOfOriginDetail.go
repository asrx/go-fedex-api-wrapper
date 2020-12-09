package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type NaftaCertificateOfOriginDetail struct {
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	BlanketPeriod *DateRange `xml:"BlanketPeriod,omitempty"`

	// Indicates which Party (if any) from the shipment is to be used as the source of importer data on the NAFTA COO form.
	ImporterSpecification *SimpleType.NaftaImporterSpecificationType `xml:"ImporterSpecification,omitempty"`

	// Contact information for "Authorized Signature" area of form.
	SignatureContact *Contact `xml:"SignatureContact,omitempty"`

	ProducerSpecification *SimpleType.NaftaProducerSpecificationType `xml:"ProducerSpecification,omitempty"`

	Producers []*NaftaProducer `xml:"Producers,omitempty"`

	CustomerImageUsages []*CustomerImageUsage `xml:"CustomerImageUsages,omitempty"`
}
