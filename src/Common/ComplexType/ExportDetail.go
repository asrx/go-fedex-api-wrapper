package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ExportDetail struct {
	B13AFilingOption *SimpleType.B13AFilingOptionType `xml:"B13AFilingOption,omitempty"`

	// General field for exporting-country-specific export data (e.g. B13A for CA, FTSR Exemption or AES Citation for US).
	ExportComplianceStatement string `xml:"ExportComplianceStatement,omitempty"`

	PermitNumber string `xml:"PermitNumber,omitempty"`

	DestinationControlDetail *DestinationControlDetail `xml:"DestinationControlDetail,omitempty"`
}
