package ComplexType

import (
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	. "github.com/asrx/go-fedex-api-wrapper/src/RateService/SimpleType"
	"time"
)

type FreightServiceCenterDetail struct {

	// Freight Industry standard non-FedEx carrier identification
	InterlineCarrierCode string `xml:"InterlineCarrierCode,omitempty"`

	// The name of the Interline carrier.
	InterlineCarrierName string `xml:"InterlineCarrierName,omitempty"`

	// Additional time it might take at the origin or destination to pickup or deliver the freight. This is usually due to the remoteness of the location. This time is included in the total transit time.
	AdditionalDays int32 `xml:"AdditionalDays,omitempty"`

	// Service branding which may be used for local pickup or delivery, distinct from service used for line-haul of customer's shipment.
	LocalService *ServiceType `xml:"LocalService,omitempty"`

	// Distance between customer address (pickup or delivery) and the supporting Freight / National Freight service center.
	LocalDistance *Distance `xml:"LocalDistance,omitempty"`

	// Time to travel between customer address (pickup or delivery) and the supporting Freight / National Freight service center.
	LocalDuration *time.Duration `xml:"LocalDuration,omitempty"`

	// Specifies when/how the customer can arrange for pickup or delivery.
	LocalServiceScheduling *FreightServiceSchedulingType `xml:"LocalServiceScheduling,omitempty"`

	// Specifies days of operation if localServiceScheduling is LIMITED.
	LimitedServiceDays []*DayOfWeekType `xml:"LimitedServiceDays,omitempty"`

	// Freight service center that is a gateway on the border of Canada or Mexico.
	GatewayLocationId string `xml:"GatewayLocationId,omitempty"`

	// Alphabetical code identifying a Freight Service Center
	Location string `xml:"Location,omitempty"`

	// Freight service center Contact and Address
	ContactAndAddress *ContactAndAddress `xml:"ContactAndAddress,omitempty"`
}
