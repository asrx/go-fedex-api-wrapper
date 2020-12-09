package ComplexType

type PackageOperationalDetail struct {

	// Human-readable text for pre-January 2011 clients.
	AstraHandlingText string `xml:"AstraHandlingText,omitempty"`

	// Human-readable content for use on a label.
	OperationalInstructions []*OperationalInstruction `xml:"OperationalInstructions,omitempty"`

	// The operational barcodes pertaining to the current package.
	Barcodes *PackageBarcodes `xml:"Barcodes,omitempty"`

	// The FedEx internal code that represents the service and/or features of service for the current package moving under a FedEx Ground service.
	GroundServiceCode string `xml:"GroundServiceCode,omitempty"`
}
