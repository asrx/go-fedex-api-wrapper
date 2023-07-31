package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	"github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"

	"github.com/asrx/go-fedex-api-wrapper/src/ShipService"
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/ShipService/ComplexType"
)

// Fedex 主单
func Test_Ship(t *testing.T) {
	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}

	var _shipper = &ComplexType.Party{
		AccountNumber: GetAccount(),
		Contact:       ShipFromContact,
		Address:       ShipFromAddrCA2,
	}

	_shipRequest := &ComplexType.RequestedShipment{
		// ShipTimestamp:     time.Now(),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		// Shipper:                       Shipper,
		Shipper:   _shipper,
		Recipient: Recipient,
		// Origin:                        nil,
		// SoldTo:                        nil,
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               SmartPostDetail,
		// BlockInsightVisibility:        false,
		LabelSpecification:        LabelSpecification,
		RateRequestTypes:          RateRequestType,
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		// TransactionDetail:       nil,
		Version:           GetVersionId("ship", 23),
		RequestedShipment: _shipRequest,
	}

	c := ShipService.NewShipPortType(_DEBUG)
	ret, err := c.ProcessShipment(processShipmentRequest)
	byte, err := json.Marshal(ret)
	fmt.Println(string(byte))
	// writeFile(byte)

	fmt.Println("-------------------")
	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	// byte, err := json.Marshal(ret)

	if err != nil || ret == nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		fmt.Println("Labe:")
		fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
	}
}

func writeFile(content []byte) error {
	file, err := os.OpenFile("ship.log", os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	defer file.Close()
	file.Write(content)
	return nil
}

// Fedex 子单
func Test_ShipSub(t *testing.T) {

	MasterTrackingId := &ComplexType.TrackingId{
		// TrackingIdType: nil,
		// FormId:         "",
		TrackingNumber: "782548213150",
	}

	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}
	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp:     time.Now(),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// Origin:                        nil,
		// SoldTo:                        nil,
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               SmartPostDetail,
		// BlockInsightVisibility:        false,
		LabelSpecification:        LabelSpecification,
		RateRequestTypes:          RateRequestType,
		MasterTrackingId:          MasterTrackingId,
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetSubPackages(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		// TransactionDetail:       nil,
		Version:           GetVersionId("ship", 23),
		RequestedShipment: _shipRequest,
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
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
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
		ShipTimestamp:          time.Now(),
		DropoffType:            &DropoffType,
		ServiceType:            &ServiceTypeSmartPost,
		PackagingType:          &PackageType,
		PreferredCurrency:      PreferredCurrency,
		Shipper:                Shipper,
		Recipient:              Recipient,
		ShippingChargesPayment: ShippingChargesPayment,
		SmartPostDetail:        SmartPostDetail,

		LabelSpecification:        LabelSpecification,
		RateRequestTypes:          RateRequestType,
		PackageCount:              &_pkgCount,
		RequestedPackageLineItems: GetPackagesSmartPost(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		// TransactionDetail:       nil,
		Version:           GetVersionId("ship", 23),
		RequestedShipment: _shipRequest,
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
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		if ret.CompletedShipmentDetail != nil {
			fmt.Println("Lable:")
			fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
		}
	}
}

// Fedex 退件单
func Test_ReturnShip(t *testing.T) {
	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}

	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp:     time.Now(),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// Origin:                        nil,
		// SoldTo:                        nil,
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               SmartPostDetail,
		// BlockInsightVisibility:        false,
		LabelSpecification:        LabelSpecification,
		RateRequestTypes:          RateRequestType,
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
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
		// TransactionDetail:       nil,
		Version:           GetVersionId("ship", 23),
		RequestedShipment: _shipRequest,
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
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		fmt.Println("Labe:")
		fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
	}
}

// Fedex express
func Test_ShipExpress(t *testing.T) {
	var diffTime = time.Hour * 24 * 9
	var _shipTimestamp = time.Now().Add(diffTime).UTC()
	fmt.Println(_shipTimestamp)
	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}

	var _shipper = &ComplexType.Party{
		AccountNumber: GetAccount(),
		Contact:       ShipFromContact,
		Address:       ShipFromAddrCA2,
	}

	_shipRequest := &ComplexType.RequestedShipment{
		ShipTimestamp: _shipTimestamp, // time.Now(),
		DropoffType:   &DropoffType,
		// ServiceType:                   &ServiceTypeStandard_overnight,
		ServiceType:       &ServiceType_EXPRESS_SAVER,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           _shipper,
		Recipient:         Recipient,
		// Origin:                        nil,
		// SoldTo:                        nil,
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               SmartPostDetail,
		// BlockInsightVisibility:        false,
		LabelSpecification:        LabelSpecification,
		RateRequestTypes:          RateRequestType,
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
	}

	processShipmentRequest := &ComplexType2.ProcessShipmentRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		// TransactionDetail:       nil,
		Version:           GetVersionId("ship", 23),
		RequestedShipment: _shipRequest,
	}

	c := ShipService.NewShipPortType(_DEBUG)
	ret, err := c.ProcessShipment(processShipmentRequest)

	if err != nil {
		t.Errorf("打单失败： %v", err)
		return
	}
	byte, err := json.Marshal(ret)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		fmt.Println("Labe:")
		if ret.CompletedShipmentDetail != nil {
			fmt.Println(string(ret.CompletedShipmentDetail.CompletedPackageDetails[0].Label.Parts[0].Image))
		}
	}
}
