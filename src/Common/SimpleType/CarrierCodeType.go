package SimpleType


// Identification of a FedEx operating company (transportation).
type CarrierCodeType string

const (
	CarrierCodeTypeFDXC CarrierCodeType = "FDXC"

	CarrierCodeTypeFDXE CarrierCodeType = "FDXE"

	CarrierCodeTypeFDXG CarrierCodeType = "FDXG"

	CarrierCodeTypeFXCC CarrierCodeType = "FXCC"

	CarrierCodeTypeFXFR CarrierCodeType = "FXFR"

	CarrierCodeTypeFXSP CarrierCodeType = "FXSP"
)