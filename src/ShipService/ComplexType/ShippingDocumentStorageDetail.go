package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"

type ShippingDocumentStorageDetail struct {

	// Indicates the mechanism by which a shipping document will be stored for later retrieval.
	Type *SimpleType.ShippingDocumentStorageDetailType `xml:"Type,omitempty"`

	// Provides the path to be used for STORED or DEFERRED documents.
	FilePath string `xml:"FilePath,omitempty"`

	// Identifies the convention by which file names are constructed for STORED or DEFERRED documents.
	FileNaming *SimpleType.ShippingDocumentNamingType `xml:"FileNaming,omitempty"`

	// Suffix to be placed at the end of the file name; required on some platforms to determine file type.
	FileSuffix string `xml:"FileSuffix,omitempty"`
}
