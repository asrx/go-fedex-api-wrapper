package SimpleType


// Specifice the kind of tax or miscellaneous charge being reported on a Commercial Invoice.
type TaxesOrMiscellaneousChargeType string

const (
	TaxesOrMiscellaneousChargeTypeCOMMISSIONS TaxesOrMiscellaneousChargeType = "COMMISSIONS"

	TaxesOrMiscellaneousChargeTypeDISCOUNTS TaxesOrMiscellaneousChargeType = "DISCOUNTS"

	TaxesOrMiscellaneousChargeTypeHANDLING_FEES TaxesOrMiscellaneousChargeType = "HANDLING_FEES"

	TaxesOrMiscellaneousChargeTypeOTHER TaxesOrMiscellaneousChargeType = "OTHER"

	TaxesOrMiscellaneousChargeTypeROYALTIES_AND_LICENSE_FEES TaxesOrMiscellaneousChargeType = "ROYALTIES_AND_LICENSE_FEES"

	TaxesOrMiscellaneousChargeTypeTAXES TaxesOrMiscellaneousChargeType = "TAXES"
)
