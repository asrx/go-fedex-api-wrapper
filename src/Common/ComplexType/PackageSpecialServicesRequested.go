package ComplexType

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type PackageSpecialServicesRequested struct {

	// The types of all special services requested for the enclosing shipment or package.
	SpecialServiceTypes []*SimpleType.PackageSpecialServiceType `xml:"SpecialServiceTypes,omitempty"`

	// For use with FedEx Ground services only; COD must be present in shipment's special services.
	CodDetail *CodDetail `xml:"CodDetail,omitempty"`

	DangerousGoodsDetail *DangerousGoodsDetail `xml:"DangerousGoodsDetail,omitempty"`

	// Provides details about the batteries or cells that are contained within this specific package.
	BatteryDetails []*BatteryClassificationDetail `xml:"BatteryDetails,omitempty"`

	DryIceWeight *Weight `xml:"DryIceWeight,omitempty"`

	SignatureOptionDetail *SignatureOptionDetail `xml:"SignatureOptionDetail,omitempty"`

	PriorityAlertDetail *PriorityAlertDetail `xml:"PriorityAlertDetail,omitempty"`

	AlcoholDetail *AlcoholDetail `xml:"AlcoholDetail,omitempty"`
}
