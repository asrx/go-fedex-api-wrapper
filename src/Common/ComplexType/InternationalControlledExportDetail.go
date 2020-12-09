package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type InternationalControlledExportDetail struct {
	Type *SimpleType.InternationalControlledExportType `xml:"Type,omitempty"`

	ForeignTradeZoneCode string `xml:"ForeignTradeZoneCode,omitempty"`

	EntryNumber string `xml:"EntryNumber,omitempty"`

	LicenseOrPermitNumber string `xml:"LicenseOrPermitNumber,omitempty"`

	LicenseOrPermitExpirationDate string `xml:"LicenseOrPermitExpirationDate,omitempty"`
}
