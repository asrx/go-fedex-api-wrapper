package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"encoding/xml"
)

type SendNotificationsReply struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/track/v16 SendNotificationsReply"`

	// This contains the severity type of the most severe Notification in the Notifications array.
	HighestSeverity *NotificationSeverityType `xml:"HighestSeverity,omitempty"`

	// Information about the request/reply such was the transaction successful or not, and any additional information relevant to the request and/or reply. There may be multiple Notifications in a reply.
	Notifications []*Notification `xml:"Notifications,omitempty"`

	// Contains the CustomerTransactionDetail that is echoed back to the caller for matching requests and replies and a Localization element for defining the language/translation used in the reply data.
	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	// Contains the version of the reply being used.
	Version *VersionId `xml:"Version,omitempty"`

	// True if duplicate packages (more than one package with the same tracking number) have been found, the packages array contains information about each duplicate. Use this information to determine which of the tracking numbers is the one you need and resend your request using the tracking number and TrackingNumberUniqueIdentifier for that package.
	DuplicateWaybill bool `xml:"DuplicateWaybill,omitempty"`

	// True if additional packages remain to be retrieved.
	MoreDataAvailable bool `xml:"MoreDataAvailable,omitempty"`

	// Value that must be passed in a TrackNotification request to retrieve the next set of packages (when MoreDataAvailable = true).
	PagingToken string `xml:"PagingToken,omitempty"`

	// Information about the notifications that are available for this tracking number. If there are duplicates the ship date and destination address information is returned for determining which TrackingNumberUniqueIdentifier to use on a subsequent request.
	Packages []*TrackNotificationPackage `xml:"Packages,omitempty"`
}
