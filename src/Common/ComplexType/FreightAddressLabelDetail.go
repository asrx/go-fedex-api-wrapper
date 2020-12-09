package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type FreightAddressLabelDetail struct {
	Format *ShippingDocumentFormat `xml:"Format,omitempty"`

	// Indicates the number of copies to be produced for each unique label.
	Copies *uint `xml:"Copies,omitempty"`

	// Specifies the quadrant of the page on which the label printing will start.
	StartingPosition *SimpleType.PageQuadrantType `xml:"StartingPosition,omitempty"`

	// If omitted, no doc tab will be produced (i.e. default = former NONE type).
	DocTabContent *DocTabContent `xml:"DocTabContent,omitempty"`

	// Controls the position of the customer specified content relative to the FedEx portion.
	CustomContentPosition *SimpleType.RelativeVerticalPositionType `xml:"CustomContentPosition,omitempty"`

	CustomContent *CustomLabelDetail `xml:"CustomContent,omitempty"`
}
