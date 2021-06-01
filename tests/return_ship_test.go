package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService"
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/ComplexType"
	"testing"
	"time"
)

func Test_Return_Ship(t *testing.T)  {
	local,_ := time.LoadLocation("Local")
	now,_ := time.ParseInLocation("2006-01-02 15:04:05","2021-06-02 14:00:00",local)

	pickupDetail := &ComplexType.PickupDetail{
		ReadyDateTime:        now,
		LatestPickupDateTime: now,
	}

	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}

	returnShipment := SimpleType.ShipmentSpecialServiceTypeRETURN_SHIPMENT
	SpecialServiceTypes := []*SimpleType.ShipmentSpecialServiceType{
		&returnShipment,
	}
	ReturnType := SimpleType.ReturnTypeFEDEX_TAG
	ReturnShipmentDetail := &ComplexType.ReturnShipmentDetail{
		ReturnType:        &ReturnType,
	}
	SpecialServicesRequested := &ComplexType.ShipmentSpecialServicesRequested{
		SpecialServiceTypes:                         SpecialServiceTypes,
		ReturnShipmentDetail:						 ReturnShipmentDetail,
	}

	_shipRequest := &ComplexType.RequestedShipment{
		DropoffType:                   &DropoffType,
		ShipTimestamp:                 now,
		ServiceType:                   &ServiceType,
		PackagingType:                 &PackageType,
		Shipper:                       Shipper,
		Recipient:                     Recipient,
		ShippingChargesPayment:        ShippingChargesPayment,
		PickupDetail: 				   pickupDetail,
		SpecialServicesRequested:	   SpecialServicesRequested,
		LabelSpecification:            LabelSpecification,
		RateRequestTypes:              RateRequestType,
		PackageCount:                  &PkgCount,
		RequestedPackageLineItems:     GetPackages(),
	}

	ProcessTagRequest := &ComplexType2.ProcessTagRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("ship", 23),
		RequestedShipment:       _shipRequest,
	}

	str,_ := json.Marshal(ProcessTagRequest)
	fmt.Println("request: ")
	fmt.Println(string(str))

	c := ShipService.NewShipPortType(false)
	ret, err := c.ProcessTag(ProcessTagRequest)

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

func Test_Cancel_Return_Ship(t *testing.T)  {
	DeleteTagRequest := &ComplexType2.DeleteTagRequest{
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

		DispatchLocationId: 	"NQAA",
		DispatchDate:           time.Now().Format("2006-01-02"),
		Payment: 				ShippingChargesPayment,
		ConfirmationNumber: 	"997091000111046",
	}

	c := ShipService.NewShipPortType(_DEBUG)
	ret, err := c.DeleteTag(DeleteTagRequest)

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