package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type RateRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/rate/v24 RateRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	// Allows the caller to specify that the transit time and commit data are to be returned in the reply.
	ReturnTransitAndCommit bool `xml:"ReturnTransitAndCommit,omitempty"`

	// Candidate carriers for rate-shopping use case. This field is only considered if requestedShipment/serviceType is omitted.
	CarrierCodes []*CarrierCodeType `xml:"CarrierCodes,omitempty"`

	// Contains zero or more service options whose combinations are to be considered when replying with available services.
	VariableOptions []*ServiceOptionType `xml:"VariableOptions,omitempty"`

	// If provided, identifies the consolidation to which this open shipment should be added after successful creation.
	ConsolidationKey *ConsolidationKey `xml:"ConsolidationKey,omitempty"`

	// The shipment for which a rate quote (or rate-shopping comparison) is desired.
	RequestedShipment *RequestedShipment `xml:"RequestedShipment,omitempty"`
}
