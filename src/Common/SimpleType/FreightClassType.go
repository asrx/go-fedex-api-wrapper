package SimpleType

// These values represent the industry-standard freight classes used for FedEx Freight and FedEx National Freight shipment description. (Note: The alphabetic prefixes are required to distinguish these values from decimal numbers on some client platforms.)
type FreightClassType string

const (
	FreightClassTypeCLASS_050 FreightClassType = "CLASS_050"

	FreightClassTypeCLASS_055 FreightClassType = "CLASS_055"

	FreightClassTypeCLASS_060 FreightClassType = "CLASS_060"

	FreightClassTypeCLASS_065 FreightClassType = "CLASS_065"

	FreightClassTypeCLASS_070 FreightClassType = "CLASS_070"

	FreightClassTypeCLASS_077_5 FreightClassType = "CLASS_077_5"

	FreightClassTypeCLASS_085 FreightClassType = "CLASS_085"

	FreightClassTypeCLASS_092_5 FreightClassType = "CLASS_092_5"

	FreightClassTypeCLASS_100 FreightClassType = "CLASS_100"

	FreightClassTypeCLASS_110 FreightClassType = "CLASS_110"

	FreightClassTypeCLASS_125 FreightClassType = "CLASS_125"

	FreightClassTypeCLASS_150 FreightClassType = "CLASS_150"

	FreightClassTypeCLASS_175 FreightClassType = "CLASS_175"

	FreightClassTypeCLASS_200 FreightClassType = "CLASS_200"

	FreightClassTypeCLASS_250 FreightClassType = "CLASS_250"

	FreightClassTypeCLASS_300 FreightClassType = "CLASS_300"

	FreightClassTypeCLASS_400 FreightClassType = "CLASS_400"

	FreightClassTypeCLASS_500 FreightClassType = "CLASS_500"
)