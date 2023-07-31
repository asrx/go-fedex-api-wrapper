package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"
	"time"

	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	rate "github.com/asrx/go-fedex-api-wrapper/src/RateService"
	"github.com/asrx/go-fedex-api-wrapper/src/RateService/ComplexType"
)

//
func TestTime(t *testing.T) {
	// 2020-11-13T06:52:49+00:00
	// fmt.Printf("%T\n", time.Now().Format(time.RFC3339))
	// fmt.Println(time.Now().UTC().Local())
	// fmt.Println(time.Now().Format(time.RFC3339))
	// 1644043113
	// fmt.Println(time.Now().Unix())
	fmt.Println(time.Now())
	fmt.Println(time.Now().UTC())

	// var tm = time.Unix(1644043,0).Format(time.RFC3339)

	if time.Unix(1654043113, 0).Before(time.Now()) {
		fmt.Println("11111")
	} else {
		fmt.Println("22222")
	}

	// var tm2 time.Time
	// fmt.Println(tm2.IsZero())
}

func Test_Rate24(t *testing.T) {
	// _SmartPostIndiciaType := rate.SmartPostIndiciaTypePARCEL_SELECT

	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp:     time.Now().UTC(), // .Format(time.RFC3339),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
		// RequestedPackageLineItems:     GetPackagesForGroup(),
	}

	// var LabelFormatType_Common2D = SimpleType.LabelFormatTypeCOMMON2D
	// var ImageType_PDF = SimpleType.ShippingDocumentImageTypePDF
	// var LabelStockTypeSTOCK_4X6 = SimpleType.LabelStockTypeSTOCK_4X6
	// labelSpecification := &ComplexType2.LabelSpecification{
	//	LabelFormatType:          &LabelFormatType_Common2D,
	//	ImageType:                &ImageType_PDF,
	//	LabelStockType:           &LabelStockTypeSTOCK_4X6,
	// }
	// _RequestedShipment.LabelSpecification = labelSpecification
	//
	// rrtList := SimpleType.RateRequestTypeLIST
	// rrtPRE := SimpleType.RateRequestTypePREFERRED
	// var RateRequestType = []*SimpleType.RateRequestType{
	//	&rrtList,
	//	&rrtPRE,
	// }
	// _RequestedShipment.RateRequestTypes = RateRequestType

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}

	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	byte, err := json.Marshal(ret)
	// byte, err := json.Marshal(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages)

	// fmt.Println(len(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages))
	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}

func Test_Rate24SmartPost(t *testing.T) {
	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp:     time.Now().UTC(), // .Format(time.RFC3339),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		SmartPostDetail:           SmartPostDetail,
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackagesSmartPost(),
	}

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}

	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	byte, err := json.Marshal(ret)
	// byte, err := json.Marshal(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages)

	fmt.Println(string(byte))
	os.Exit(1)

	fmt.Println(len(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages))
	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}

func Test_RateHomeDelivery24(t *testing.T) {
	// _SmartPostIndiciaType := rate.SmartPostIndiciaTypePARCEL_SELECT

	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp:     time.Now().UTC(), // .Format(time.RFC3339),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceTypeHomeDelivery,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
	}

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}

	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

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

func Test_RateReturnTag24(t *testing.T) {
	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp:     time.Now().UTC(), // .Format(time.RFC3339),
		DropoffType:       &DropoffType,
		ServiceType:       &ServiceType,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
		// RequestedPackageLineItems:     GetPackagesForGroup(),
	}

	// var LabelFormatType_Common2D = SimpleType.LabelFormatTypeCOMMON2D
	// var ImageType_PDF = SimpleType.ShippingDocumentImageTypePDF
	// var LabelStockTypeSTOCK_4X6 = SimpleType.LabelStockTypeSTOCK_4X6
	// labelSpecification := &ComplexType2.LabelSpecification{
	//	LabelFormatType:          &LabelFormatType_Common2D,
	//	ImageType:                &ImageType_PDF,
	//	LabelStockType:           &LabelStockTypeSTOCK_4X6,
	// }
	// _RequestedShipment.LabelSpecification = labelSpecification
	//
	// rrtList := SimpleType.RateRequestTypeLIST
	// rrtPRE := SimpleType.RateRequestTypePREFERRED
	// var RateRequestType = []*SimpleType.RateRequestType{
	//	&rrtList,
	//	&rrtPRE,
	// }
	// _RequestedShipment.RateRequestTypes = RateRequestType

	// 退件参数
	returnType := SimpleType.ReturnTypePRINT_RETURN_LABEL
	shipmentSpecialServiceType := SimpleType.ShipmentSpecialServiceTypeRETURN_SHIPMENT
	_RequestedShipment.SpecialServicesRequested = &ComplexType2.ShipmentSpecialServicesRequested{
		SpecialServiceTypes: []*SimpleType.ShipmentSpecialServiceType{
			&shipmentSpecialServiceType,
		},
		ReturnShipmentDetail: &ComplexType2.ReturnShipmentDetail{
			ReturnType: &returnType,
		},
	}

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}
	byte, err := json.Marshal(_RateRequest)
	fmt.Println(string(byte))
	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	byte, err = json.Marshal(ret)
	// byte, err := json.Marshal(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages)

	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}

func Test_RateStandardOvernight(t *testing.T) {
	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp: time.Now().UTC(), // .Format(time.RFC3339),
		DropoffType:   &DropoffType,
		// Overnight
		ServiceType:       &ServiceTypeStandard_overnight,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
		// RequestedPackageLineItems:     GetPackagesForGroup(),
	}

	// var LabelFormatType_Common2D = SimpleType.LabelFormatTypeCOMMON2D
	// var ImageType_PDF = SimpleType.ShippingDocumentImageTypePDF
	// var LabelStockTypeSTOCK_4X6 = SimpleType.LabelStockTypeSTOCK_4X6
	// labelSpecification := &ComplexType2.LabelSpecification{
	//	LabelFormatType:          &LabelFormatType_Common2D,
	//	ImageType:                &ImageType_PDF,
	//	LabelStockType:           &LabelStockTypeSTOCK_4X6,
	// }
	// _RequestedShipment.LabelSpecification = labelSpecification

	rrtList := SimpleType.RateRequestTypeLIST
	rrtPRE := SimpleType.RateRequestTypePREFERRED
	var RateRequestType = []*SimpleType.RateRequestType{
		&rrtList,
		&rrtPRE,
	}
	_RequestedShipment.RateRequestTypes = RateRequestType

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}

	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	req, err := json.Marshal(_RateRequest)
	fmt.Println(string(req))
	byte, err := json.Marshal(ret)
	// byte, err := json.Marshal(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages)

	// fmt.Println(len(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages))
	if err != nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		fmt.Println("================")
		fmt.Println(ret.RateReplyDetails[0].RatedShipmentDetails[0].ShipmentRateDetail.TotalNetCharge)
	}
}

func Test_Datetime(t *testing.T) {
	// var _shipTimestamp = time.Now().Add(time.Duration(100))
	var diffTime = time.Hour * 24 * 11
	fmt.Println(time.Duration(diffTime))
	fmt.Println(time.Now().Add(diffTime).UTC().Format(time.RFC3339))
}

func Test_RatePriorityOvernight(t *testing.T) {

	var diffTime = time.Hour * 24 * 9
	// fmt.Println(time.Duration(diffTime))
	// fmt.Println(time.Now().Add(diffTime).UTC().Format(time.RFC3339))
	var _shipTimestamp = time.Now().Add(diffTime).UTC()
	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp: _shipTimestamp, // time.Now().UTC(),//.Format(time.RFC3339),
		DropoffType:   &DropoffType,
		// Overnight
		ServiceType:       &ServiceTypePriority_overnight,
		PackagingType:     &PackageType,
		PreferredCurrency: PreferredCurrency,
		Shipper:           Shipper,
		Recipient:         Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		// Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		// SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment: ShippingChargesPayment,
		// SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		// },
		PackageCount:              &PkgCount,
		RequestedPackageLineItems: GetPackages(),
		// RequestedPackageLineItems:     GetPackagesForGroup(),
	}

	// var LabelFormatType_Common2D = SimpleType.LabelFormatTypeCOMMON2D
	// var ImageType_PDF = SimpleType.ShippingDocumentImageTypePDF
	// var LabelStockTypeSTOCK_4X6 = SimpleType.LabelStockTypeSTOCK_4X6
	// labelSpecification := &ComplexType2.LabelSpecification{
	//	LabelFormatType:          &LabelFormatType_Common2D,
	//	ImageType:                &ImageType_PDF,
	//	LabelStockType:           &LabelStockTypeSTOCK_4X6,
	// }
	// _RequestedShipment.LabelSpecification = labelSpecification
	//
	// rrtList := SimpleType.RateRequestTypeLIST
	// rrtPRE := SimpleType.RateRequestTypePREFERRED
	// var RateRequestType = []*SimpleType.RateRequestType{
	//	&rrtList,
	//	&rrtPRE,
	// }
	// _RequestedShipment.RateRequestTypes = RateRequestType

	_versionId := GetVersionId("crs", 24)
	_RateRequest := &ComplexType.RateRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		ReturnTransitAndCommit:  false,
		RequestedShipment:       _RequestedShipment,
	}

	c := rate.NewRatePortType(_DEBUG)
	ret, err := c.GetRates(_RateRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
	}
	req, err := json.Marshal(_RateRequest)
	fmt.Println(string(req))
	byte, err := json.Marshal(ret)
	// byte, err := json.Marshal(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages)

	// fmt.Println(len(ret.RateReplyDetails[0].RatedShipmentDetails[0].RatedPackages))
	if err != nil || ret == nil {
		fmt.Println("Marshal Error: ")
		fmt.Printf("%+v\n", ret)
	} else {
		fmt.Printf("%+v\n", ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
		fmt.Println("================")
		fmt.Println(ret.RateReplyDetails[0].RatedShipmentDetails[0].ShipmentRateDetail.TotalNetCharge)
	}
}
