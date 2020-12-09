package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type ClientDetail struct {

	// The FedEx account number associated with this transaction.
	AccountNumber string `xml:"AccountNumber,omitempty"`

	// This number is assigned by FedEx and identifies the unique device from which the request is originating
	MeterNumber string `xml:"MeterNumber,omitempty"`

	// Only used in transactions which require identification of the FedEx Office integrator.
	IntegratorId string `xml:"IntegratorId,omitempty"`

	// Indicates the region from which the transaction is submitted.
	Region *SimpleType.ExpressRegionCode `xml:"Region,omitempty"`

	// The language to be used for human-readable Notification.localizedMessages in responses to the request containing this ClientDetail object. Different requests from the same client may contain different Localization data. (Contrast with TransactionDetail.localization, which governs data payload language/translation.)
	Localization *Localization `xml:"Localization,omitempty"`
}
