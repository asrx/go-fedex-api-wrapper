package tests

import (
	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"time"

	"encoding/json"
	"fmt"
	"github.com/asrx/go-fedex-api-wrapper/src/ShipService"
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/ComplexType"
	"testing"
)

// Fedex 主单
func Test_Ship(t *testing.T) {
	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}
	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp:                 time.Now(),
		DropoffType:                   &DropoffType,
		ServiceType:                   &ServiceType,
		PackagingType:                 &PackageType,
		PreferredCurrency:             PreferredCurrency,
		Shipper:                       Shipper,
		Recipient:                     Recipient,
		//Origin:                        nil,
		//SoldTo:                        nil,
		ShippingChargesPayment:        ShippingChargesPayment,
		//SmartPostDetail:               SmartPostDetail,
		//BlockInsightVisibility:        false,
		LabelSpecification:            LabelSpecification,
		RateRequestTypes:              RateRequestType,
		PackageCount:                  &PkgCount,
		RequestedPackageLineItems:     GetPackages(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("ship", 23),
		RequestedShipment:       _shipRequest,
	}

	c := ShipService.NewShipPortType(_DEBUG)
	ret, err := c.ProcessShipment(processShipmentRequest)

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

// Fedex 子单
func Test_ShipSub(t *testing.T) {

	MasterTrackingId := &ComplexType.TrackingId{
		//TrackingIdType: nil,
		//FormId:         "",
		TrackingNumber: "782548213150",
	}

	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}
	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp:                 time.Now(),
		DropoffType:                   &DropoffType,
		ServiceType:                   &ServiceType,
		PackagingType:                 &PackageType,
		PreferredCurrency:             PreferredCurrency,
		Shipper:                       Shipper,
		Recipient:                     Recipient,
		//Origin:                        nil,
		//SoldTo:                        nil,
		ShippingChargesPayment:        ShippingChargesPayment,
		//SmartPostDetail:               SmartPostDetail,
		//BlockInsightVisibility:        false,
		LabelSpecification:            LabelSpecification,
		RateRequestTypes:              RateRequestType,
		MasterTrackingId: 			   MasterTrackingId,
		PackageCount:                  &PkgCount,
		RequestedPackageLineItems:     GetPackages2(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("ship", 23),
		RequestedShipment:       _shipRequest,
	}

	c := ShipService.NewShipPortType(false)
	ret, err := c.ProcessShipment(processShipmentRequest)

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

// Fedex SmartPost
func Test_ShipSmartPost(t *testing.T) {
	fmt.Println("======= Fedex SmartPost =======")
	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}
	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp:                 time.Now(),
		DropoffType:                   &DropoffType,
		ServiceType:                   &ServiceTypeSmartPost,
		PackagingType:                 &PackageType,
		PreferredCurrency:             PreferredCurrency,
		Shipper:                       Shipper,
		Recipient:                     Recipient,
		ShippingChargesPayment:        ShippingChargesPayment,
		SmartPostDetail:               SmartPostDetail,

		LabelSpecification:            LabelSpecification,
		RateRequestTypes:              RateRequestType,
		PackageCount:                  &PkgCount,
		RequestedPackageLineItems:     GetPackagesSmartPost(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("ship", 23),
		RequestedShipment:       _shipRequest,
	}

	c := ShipService.NewShipPortType(false)
	ret, err := c.ProcessShipment(processShipmentRequest)

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