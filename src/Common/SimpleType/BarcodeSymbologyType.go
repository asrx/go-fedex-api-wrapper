package SimpleType


// Identification of the type of barcode (symbology) used on FedEx documents and labels.
type BarcodeSymbologyType string

const (
	BarcodeSymbologyTypeCODABAR BarcodeSymbologyType = "CODABAR"

	BarcodeSymbologyTypeCODE128 BarcodeSymbologyType = "CODE128"

	BarcodeSymbologyTypeCODE128B BarcodeSymbologyType = "CODE128B"

	BarcodeSymbologyTypeCODE128C BarcodeSymbologyType = "CODE128C"

	BarcodeSymbologyTypeCODE128_WIDEBAR BarcodeSymbologyType = "CODE128_WIDEBAR"

	BarcodeSymbologyTypeCODE39 BarcodeSymbologyType = "CODE39"

	BarcodeSymbologyTypeCODE93 BarcodeSymbologyType = "CODE93"

	BarcodeSymbologyTypeI2OF5 BarcodeSymbologyType = "I2OF5"

	BarcodeSymbologyTypeMANUAL BarcodeSymbologyType = "MANUAL"

	BarcodeSymbologyTypePDF417 BarcodeSymbologyType = "PDF417"

	BarcodeSymbologyTypePOSTNET BarcodeSymbologyType = "POSTNET"

	BarcodeSymbologyTypeQR_CODE BarcodeSymbologyType = "QR_CODE"

	BarcodeSymbologyTypeUCC128 BarcodeSymbologyType = "UCC128"
)
