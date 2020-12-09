package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type SmartPostShipmentDetail struct {
	ProcessingOptionsRequested *SmartPostShipmentProcessingOptionsRequested `xml:"ProcessingOptionsRequested,omitempty"`

	Indicia *SimpleType.SmartPostIndiciaType `xml:"Indicia,omitempty"`

	AncillaryEndorsement *SimpleType.SmartPostAncillaryEndorsementType `xml:"AncillaryEndorsement,omitempty"`

	HubId string `xml:"HubId,omitempty"`

	CustomerManifestId string `xml:"CustomerManifestId,omitempty"`
}
