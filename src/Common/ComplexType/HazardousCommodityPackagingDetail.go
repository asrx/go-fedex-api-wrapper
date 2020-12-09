package ComplexType

type HazardousCommodityPackagingDetail struct {

	// Number of units of the type below.
	Count *uint `xml:"Count,omitempty"`

	// Units in which the hazardous commodity is packaged.
	Units string `xml:"Units,omitempty"`
}
