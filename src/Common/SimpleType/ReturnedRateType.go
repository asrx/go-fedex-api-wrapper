package SimpleType


// The "PAYOR..." rates are expressed in the currency identified in the payor's rate table(s). The "RATED..." rates are expressed in the currency of the origin country. Former "...COUNTER..." values have become "...RETAIL..." values, except for PAYOR_COUNTER and RATED_COUNTER, which have been removed.
type ReturnedRateType string

const (
	ReturnedRateTypeNEGOTIATED ReturnedRateType = "NEGOTIATED"

	ReturnedRateTypePAYOR_ACCOUNT_PACKAGE ReturnedRateType = "PAYOR_ACCOUNT_PACKAGE"

	ReturnedRateTypePAYOR_ACCOUNT_SHIPMENT ReturnedRateType = "PAYOR_ACCOUNT_SHIPMENT"

	ReturnedRateTypePAYOR_LIST_PACKAGE ReturnedRateType = "PAYOR_LIST_PACKAGE"

	ReturnedRateTypePAYOR_LIST_SHIPMENT ReturnedRateType = "PAYOR_LIST_SHIPMENT"

	ReturnedRateTypePREFERRED_ACCOUNT_PACKAGE ReturnedRateType = "PREFERRED_ACCOUNT_PACKAGE"

	ReturnedRateTypePREFERRED_ACCOUNT_SHIPMENT ReturnedRateType = "PREFERRED_ACCOUNT_SHIPMENT"

	ReturnedRateTypePREFERRED_LIST_PACKAGE ReturnedRateType = "PREFERRED_LIST_PACKAGE"

	ReturnedRateTypePREFERRED_LIST_SHIPMENT ReturnedRateType = "PREFERRED_LIST_SHIPMENT"

	ReturnedRateTypePREFERRED_NEGOTIATED ReturnedRateType = "PREFERRED_NEGOTIATED"
)

