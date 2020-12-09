package SimpleType


// Specifies the type of deletion to be performed on a shipment.
type DeletionControlType string

const (
	DeletionControlTypeDELETE_ALL_PACKAGES DeletionControlType = "DELETE_ALL_PACKAGES"

	DeletionControlTypeDELETE_ENTIRE_CONSOLIDATION DeletionControlType = "DELETE_ENTIRE_CONSOLIDATION"

	DeletionControlTypeDELETE_ONE_PACKAGE DeletionControlType = "DELETE_ONE_PACKAGE"

	DeletionControlTypeLEGACY DeletionControlType = "LEGACY"
)

