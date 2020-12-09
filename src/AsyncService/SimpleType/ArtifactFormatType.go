package SimpleType

// Identifies the format of the artifact.
type ArtifactFormatType string

const (
	ArtifactFormatTypeBINARY ArtifactFormatType = "BINARY"

	ArtifactFormatTypeDOC ArtifactFormatType = "DOC"

	ArtifactFormatTypeEPL2 ArtifactFormatType = "EPL2"

	ArtifactFormatTypePDF ArtifactFormatType = "PDF"

	ArtifactFormatTypePNG ArtifactFormatType = "PNG"

	ArtifactFormatTypeRTF ArtifactFormatType = "RTF"

	ArtifactFormatTypeTEXT ArtifactFormatType = "TEXT"

	ArtifactFormatTypeXML ArtifactFormatType = "XML"

	ArtifactFormatTypeZPLII ArtifactFormatType = "ZPLII"
)
