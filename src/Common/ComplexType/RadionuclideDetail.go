package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type RadionuclideDetail struct {
	Radionuclide string `xml:"Radionuclide,omitempty"`

	Activity *RadionuclideActivity `xml:"Activity,omitempty"`

	// Indicates whether packaging type "EXCEPTED" or "EXCEPTED_PACKAGE" is for radioactive material in reportable quantity.
	ExceptedPackagingIsReportableQuantity bool `xml:"ExceptedPackagingIsReportableQuantity,omitempty"`

	PhysicalForm *SimpleType.PhysicalFormType `xml:"PhysicalForm,omitempty"`

	ChemicalForm string `xml:"ChemicalForm,omitempty"`
}
