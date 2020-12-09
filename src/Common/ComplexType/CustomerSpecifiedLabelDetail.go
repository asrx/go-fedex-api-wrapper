package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CustomerSpecifiedLabelDetail struct {

	// If omitted, no doc tab will be produced (i.e. default is former NONE type).
	DocTabContent *DocTabContent `xml:"DocTabContent,omitempty"`

	// Controls the position of the customer specified content relative to the FedEx portion.
	CustomContentPosition *SimpleType.RelativeVerticalPositionType `xml:"CustomContentPosition,omitempty"`

	CustomContent *CustomLabelDetail `xml:"CustomContent,omitempty"`

	ConfigurableReferenceEntries []*ConfigurableLabelReferenceEntry `xml:"ConfigurableReferenceEntries,omitempty"`

	// Controls which data/sections will be suppressed.
	MaskedData []*SimpleType.LabelMaskableDataType `xml:"MaskedData,omitempty"`

	// For customers producing their own Ground labels, this field specifies which secondary barcode will be printed on the label; so that the primary barcode produced by FedEx has the correct SCNC.
	SecondaryBarcode *SimpleType.SecondaryBarcodeType `xml:"SecondaryBarcode,omitempty"`

	TermsAndConditionsLocalization *Localization `xml:"TermsAndConditionsLocalization,omitempty"`

	RegulatoryLabels []*RegulatoryLabelContentDetail `xml:"RegulatoryLabels,omitempty"`

	// Controls the number of additional copies of supplemental labels.
	AdditionalLabels []*AdditionalLabelsDetail `xml:"AdditionalLabels,omitempty"`

	// This value reduces the default quantity of destination/consignee air waybill labels. A value of zero indicates no change to default. A minimum of one copy will always be produced.
	AirWaybillSuppressionCount *uint `xml:"AirWaybillSuppressionCount,omitempty"`
}
