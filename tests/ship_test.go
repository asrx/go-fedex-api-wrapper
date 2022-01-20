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
		fmt.Println("Labe:")
		fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
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
		RequestedPackageLineItems:     GetSubPackages(),
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
	var _pkgCount uint = 1
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
		PackageCount:                  &_pkgCount,
		RequestedPackageLineItems:     GetPackagesSmartPost(),
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
		if ret.CompletedShipmentDetail != nil {
			fmt.Println("Lable:")
			fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
		}
	}
}

// Fedex 退件单
func Test_ReturnShip(t *testing.T){
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

	// 退件参数
	returnType := SimpleType.ReturnTypePRINT_RETURN_LABEL
	shipmentSpecialServiceType := SimpleType.ShipmentSpecialServiceTypeRETURN_SHIPMENT
	_shipRequest.SpecialServicesRequested = &ComplexType.ShipmentSpecialServicesRequested{
		SpecialServiceTypes: []*SimpleType.ShipmentSpecialServiceType{
			&shipmentSpecialServiceType,
		},
		ReturnShipmentDetail: &ComplexType.ReturnShipmentDetail{
			ReturnType: &returnType,
		},
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
		fmt.Println("Labe:")
		fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
	}
}