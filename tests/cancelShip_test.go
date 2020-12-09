package tests

import (
	"encoding/json"
	"fmt"
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService/ComplexType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/SimpleType"
	"testing"
	"time"
)

func Test_CancelShip(t *testing.T) {
	_trackIdType := SimpleType.TrackingIdTypeFEDEX

	_deletionControl := SimpleType2.DeletionControlTypeDELETE_ALL_PACKAGES

	//var i1 int32 = 1
	deleteShipmentRequest := &ComplexType.DeleteShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("ship", 23),
		//	&ComplexType2.VersionId{
		//	ServiceId:    "ship",
		//	Major:        12,
		//	Intermediate: &i1,
		//	Minor:        &I32,
		//},
		ShipTimestamp:           time.Now(),
		TrackingId:              &ComplexType2.TrackingId{
			TrackingIdType: &_trackIdType,
			TrackingNumber: "399541845233",
		},
		DeletionControl:         &_deletionControl,
	}

	c := ShipService.NewShipPortType(true)
	ret, err := c.DeleteShipment(deleteShipmentRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	byte, err := json.Marshal(ret)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n",ret)
	}else{
		fmt.Printf("%+v\n",ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}
