package ComplexType

type RadioactivityDetail struct {
	TransportIndex float64 `xml:"TransportIndex,omitempty"`

	SurfaceReading float64 `xml:"SurfaceReading,omitempty"`

	CriticalitySafetyIndex float64 `xml:"CriticalitySafetyIndex,omitempty"`

	Dimensions *Dimensions `xml:"Dimensions,omitempty"`
}
