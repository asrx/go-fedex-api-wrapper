package SimpleType

// Specifies the way in which base charges for a Freight shipment or shipment leg are calculated.
type FreightBaseChargeCalculationType string

const (
	FreightBaseChargeCalculationTypeLINE_ITEMS FreightBaseChargeCalculationType = "LINE_ITEMS"

	FreightBaseChargeCalculationTypeUNIT_PRICING FreightBaseChargeCalculationType = "UNIT_PRICING"
)