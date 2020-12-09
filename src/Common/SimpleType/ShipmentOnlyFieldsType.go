package SimpleType


// These values identify which package-level data values will be provided at the shipment-level.
type ShipmentOnlyFieldsType string

const (
	ShipmentOnlyFieldsTypeDIMENSIONS ShipmentOnlyFieldsType = "DIMENSIONS"

	ShipmentOnlyFieldsTypeINSURED_VALUE ShipmentOnlyFieldsType = "INSURED_VALUE"

	ShipmentOnlyFieldsTypeWEIGHT ShipmentOnlyFieldsType = "WEIGHT"
)
