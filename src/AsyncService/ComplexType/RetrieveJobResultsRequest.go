package ComplexType

import (
	"encoding/xml"
	. "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
)

type RetrieveJobResultsRequest struct {
	XMLName xml.Name `xml:"http://fedex.com/ws/async/v4 RetrieveJobResultsRequest"`

	// Descriptive data to be used in authentication of the sender's identity (and right to use FedEx web services).
	WebAuthenticationDetail *WebAuthenticationDetail `xml:"WebAuthenticationDetail,omitempty"`

	ClientDetail *ClientDetail `xml:"ClientDetail,omitempty"`

	TransactionDetail *TransactionDetail `xml:"TransactionDetail,omitempty"`

	Version *VersionId `xml:"Version,omitempty"`

	// Specifies the job under which the desired artifacts are stored.
	JobId string `xml:"JobId,omitempty"`

	// Specifies the filters to be used for retrieving artifacts.
	Filters []*ArtifactRetrievalFilter `xml:"Filters,omitempty"`
}
