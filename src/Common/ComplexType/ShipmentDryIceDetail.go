package ComplexType

type ShipmentDryIceDetail struct {

	// Total number of packages in the shipment that contain dry ice.
	PackageCount *uint `xml:"PackageCount,omitempty"`

	// Total shipment dry ice weight for all packages.
	TotalWeight *Weight `xml:"TotalWeight,omitempty"`

	ProcessingOptions *ShipmentDryIceProcessingOptionsRequested `xml:"ProcessingOptions,omitempty"`
}
