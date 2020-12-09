package ComplexType


type VersionId struct {

	// Identifies a system or sub-system which performs an operation.
	ServiceId string `xml:"ServiceId,omitempty"`

	// Identifies the service business level.
	Major int32 `xml:"Major,omitempty"`

	// Identifies the service interface level.
	Intermediate *int32 `xml:"Intermediate,omitempty"`

	// Identifies the service code level.
	Minor *int32 `xml:"Minor,omitempty"`
}

