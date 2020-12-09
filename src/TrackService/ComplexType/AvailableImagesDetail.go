package ComplexType

import . "github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"

type AvailableImagesDetail struct {
	Type *AvailableImageType `xml:"Type,omitempty"`

	Size *ImageSizeType `xml:"Size,omitempty"`
}
