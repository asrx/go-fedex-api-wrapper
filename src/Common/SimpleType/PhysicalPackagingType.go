package SimpleType

// This enumeration rationalizes the former FedEx Express international "admissibility package" types (based on ANSI X.12) and the FedEx Freight packaging types. The values represented are those common to both carriers.
type PhysicalPackagingType string

const (
	PhysicalPackagingTypeBAG PhysicalPackagingType = "BAG"

	PhysicalPackagingTypeBARREL PhysicalPackagingType = "BARREL"

	PhysicalPackagingTypeBASKET PhysicalPackagingType = "BASKET"

	PhysicalPackagingTypeBOX PhysicalPackagingType = "BOX"

	PhysicalPackagingTypeBUCKET PhysicalPackagingType = "BUCKET"

	PhysicalPackagingTypeBUNDLE PhysicalPackagingType = "BUNDLE"

	PhysicalPackagingTypeCARTON PhysicalPackagingType = "CARTON"

	PhysicalPackagingTypeCASE PhysicalPackagingType = "CASE"

	PhysicalPackagingTypeCONTAINER PhysicalPackagingType = "CONTAINER"

	PhysicalPackagingTypeCRATE PhysicalPackagingType = "CRATE"

	PhysicalPackagingTypeCYLINDER PhysicalPackagingType = "CYLINDER"

	PhysicalPackagingTypeDRUM PhysicalPackagingType = "DRUM"

	PhysicalPackagingTypeENVELOPE PhysicalPackagingType = "ENVELOPE"

	PhysicalPackagingTypeHAMPER PhysicalPackagingType = "HAMPER"

	PhysicalPackagingTypeOTHER PhysicalPackagingType = "OTHER"

	PhysicalPackagingTypePAIL PhysicalPackagingType = "PAIL"

	PhysicalPackagingTypePALLET PhysicalPackagingType = "PALLET"

	PhysicalPackagingTypePIECE PhysicalPackagingType = "PIECE"

	PhysicalPackagingTypeREEL PhysicalPackagingType = "REEL"

	PhysicalPackagingTypeROLL PhysicalPackagingType = "ROLL"

	PhysicalPackagingTypeSKID PhysicalPackagingType = "SKID"

	PhysicalPackagingTypeTANK PhysicalPackagingType = "TANK"

	PhysicalPackagingTypeTUBE PhysicalPackagingType = "TUBE"
)