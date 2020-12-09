package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/TrackService/SimpleType"
	"testing"
)

func Test_TrackByNo(t *testing.T)  {
	_trackingId1 := "396716838866"

	_TrackIdentifierType := SimpleType.TrackIdentifierTypeTRACKING_NUMBER_OR_DOORTAG
	// detail1
	track1 := ComplexType.TrackSelectionDetail{
		PackageIdentifier:              &ComplexType.TrackPackageIdentifier{
			Type:  &_TrackIdentifierType,
			Value: _trackingId1,
		},
	}

	_details := []*ComplexType.TrackSelectionDetail{
		&track1,
	}

	_processingOption := SimpleType.TrackRequestProcessingOptionTypeINCLUDE_DETAILED_SCANS
	_trackRequet := &ComplexType.TrackRequest{
		WebAuthenticationDetail:               GetWebAuthenticationDetail(),
		ClientDetail:                          GetClientDetail(),
		Version:                               GetVersionId("trck",16),
		SelectionDetails:                      _details,
		ProcessingOptions: 					   []*SimpleType.TrackRequestProcessingOptionType{ &_processingOption },
	}

	c := TrackService.NewTrackPortType(false)
	ret, err := c.Track(_trackRequet)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}

	// Print TrackEvent
	for _, completedTrackDetail := range ret.CompletedTrackDetails {
		for _, details := range completedTrackDetail.TrackDetails {
			for _, event := range details.Events{
				fmt.Printf("%s | %s\n", event.Timestamp, event.EventDescription)
			}
		}
	}

	byte, err := json.Marshal(ret)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n",ret)
	}else{
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}
