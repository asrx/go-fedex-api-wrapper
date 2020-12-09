package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

type CompletedTrackDetail struct {
	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	Notifications []*Notification `xml:"Notifications,omitempty"`

	// True if duplicate packages (more than one package with the same tracking number) have been found, and only limited data will be provided for each one.
	DuplicateWaybill bool `xml:"DuplicateWaybill,omitempty"`

	// True if additional packages remain to be retrieved.
	MoreData bool `xml:"MoreData,omitempty"`

	// Value that must be passed in a TrackNotification request to retrieve the next set of packages (when MoreDataAvailable = true).
	PagingToken string `xml:"PagingToken,omitempty"`

	// Identifies the total number of available track details across all pages.
	TrackDetailsCount *uint `xml:"TrackDetailsCount,omitempty"`

	// Contains detailed tracking information for the requested packages(s).
	TrackDetails []*TrackDetail `xml:"TrackDetails,omitempty"`
}
