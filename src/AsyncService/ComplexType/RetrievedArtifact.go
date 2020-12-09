package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/AsyncService/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type RetrievedArtifact struct {
	AccessReference string `xml:"AccessReference,omitempty"`

	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	// Identifies the type of artifact.
	Type *ArtifactType `xml:"Type,omitempty"`

	// Specifies the format of the artifact.
	Format *ArtifactFormatType `xml:"Format,omitempty"`

	// Identifies the schema or specific format definition used for the artifact.
	FormatSpecification string `xml:"FormatSpecification,omitempty"`

	// Specifies a reference to the artifact that follows the naming convention defined for file storage.
	ReferenceId string `xml:"ReferenceId,omitempty"`

	Parts []*ArtifactPart `xml:"Parts,omitempty"`
}
