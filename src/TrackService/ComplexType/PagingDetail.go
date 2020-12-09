package ComplexType

type PagingDetail struct {

	// When the MoreData field = true in a TrackReply the PagingToken must be sent in the subsequent TrackRequest to retrieve the next page of data.
	PagingToken string `xml:"PagingToken,omitempty"`

	// Specifies the number of results to display per page when the there is more than one page in the subsequent TrackReply.
	NumberOfResultsPerPage *uint `xml:"NumberOfResultsPerPage,omitempty"`
}
