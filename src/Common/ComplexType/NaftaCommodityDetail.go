package ComplexType

import "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

type NaftaCommodityDetail struct {

	// Defined by NAFTA regulations.
	PreferenceCriterion *SimpleType.NaftaPreferenceCriterionCode `xml:"PreferenceCriterion,omitempty"`

	// Defined by NAFTA regulations.
	ProducerDetermination *SimpleType.NaftaProducerDeterminationCode `xml:"ProducerDetermination,omitempty"`

	// Identification of which producer is associated with this commodity (if multiple producers are used in a single shipment).
	ProducerId string `xml:"ProducerId,omitempty"`

	NetCostMethod *SimpleType.NaftaNetCostMethodCode `xml:"NetCostMethod,omitempty"`

	// Date range over which RVC net cost was calculated.
	NetCostDateRange *DateRange `xml:"NetCostDateRange,omitempty"`
}
