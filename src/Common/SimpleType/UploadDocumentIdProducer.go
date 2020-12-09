package SimpleType


// Specifies the application that is responsible for managing the document id.
type UploadDocumentIdProducer string

const (
	UploadDocumentIdProducerCUSTOMER UploadDocumentIdProducer = "CUSTOMER"

	UploadDocumentIdProducerFEDEX_CAFE UploadDocumentIdProducer = "FEDEX_CAFE"

	UploadDocumentIdProducerFEDEX_CSHP UploadDocumentIdProducer = "FEDEX_CSHP"

	UploadDocumentIdProducerFEDEX_FXRS UploadDocumentIdProducer = "FEDEX_FXRS"

	UploadDocumentIdProducerFEDEX_GSMW UploadDocumentIdProducer = "FEDEX_GSMW"

	UploadDocumentIdProducerFEDEX_GTM UploadDocumentIdProducer = "FEDEX_GTM"

	UploadDocumentIdProducerFEDEX_INET UploadDocumentIdProducer = "FEDEX_INET"
)
