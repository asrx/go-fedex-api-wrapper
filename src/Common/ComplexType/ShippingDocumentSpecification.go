package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ShippingDocumentSpecification struct {

	// Indicates the types of shipping documents requested by the shipper.
	ShippingDocumentTypes []*SimpleType.RequestedShippingDocumentType `xml:"ShippingDocumentTypes,omitempty"`

	CertificateOfOrigin *CertificateOfOriginDetail `xml:"CertificateOfOrigin,omitempty"`

	CommercialInvoiceDetail *CommercialInvoiceDetail `xml:"CommercialInvoiceDetail,omitempty"`

	// Specifies the production of each package-level custom document (the same specification is used for all packages).
	CustomPackageDocumentDetail []*CustomDocumentDetail `xml:"CustomPackageDocumentDetail,omitempty"`

	// Specifies the production of a shipment-level custom document.
	CustomShipmentDocumentDetail []*CustomDocumentDetail `xml:"CustomShipmentDocumentDetail,omitempty"`

	ExportDeclarationDetail *ExportDeclarationDetail `xml:"ExportDeclarationDetail,omitempty"`

	GeneralAgencyAgreementDetail *GeneralAgencyAgreementDetail `xml:"GeneralAgencyAgreementDetail,omitempty"`

	NaftaCertificateOfOriginDetail *NaftaCertificateOfOriginDetail `xml:"NaftaCertificateOfOriginDetail,omitempty"`

	// Specifies the production of the OP-900 document for hazardous materials packages.
	Op900Detail *Op900Detail `xml:"Op900Detail,omitempty"`

	// Specifies the production of the 1421c document for dangerous goods shipment.
	DangerousGoodsShippersDeclarationDetail *DangerousGoodsShippersDeclarationDetail `xml:"DangerousGoodsShippersDeclarationDetail,omitempty"`

	// Specifies the production of the OP-900 document for hazardous materials.
	FreightAddressLabelDetail *FreightAddressLabelDetail `xml:"FreightAddressLabelDetail,omitempty"`

	// Specifies the production of the return instructions document.
	ReturnInstructionsDetail *ReturnInstructionsDetail `xml:"ReturnInstructionsDetail,omitempty"`
}
