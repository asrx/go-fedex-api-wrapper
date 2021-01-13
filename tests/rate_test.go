package tests

import (
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	rate "github.com/asrx/go-fedex-api-wrapper/src/RateService"
	"github.com/asrx/go-fedex-api-wrapper/src/RateService/ComplexType"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)


//
func TestTime(t *testing.T)  {
	//2020-11-13T06:52:49+00:00
	fmt.Printf("%T\n", time.Now().Format(time.RFC3339))
	fmt.Println(time.Now().UTC().Local())
}

func Test_Rate24(t *testing.T){
	//_SmartPostIndiciaType := rate.SmartPostIndiciaTypePARCEL_SELECT

	_RequestedShipment := &ComplexType2.RequestedShipment{
		ShipTimestamp:                 time.Now().UTC(),//.Format(time.RFC3339),
		DropoffType:                   &DropoffType,
		ServiceType:                   &ServiceType,
		PackagingType:                 &PackageType,
		PreferredCurrency:             PreferredCurrency,
		Shipper:                       Shipper,
		Recipient:                     Recipient,
		// 货物的实际起始地址，如果与发货人的地址不同, 则使用。
		//Origin:                        nil,
		// 被售方用于清关;例如，支持美国进口海关规则。对该字段的需求可能因是否在合并中指定了卖方而有所不同。
		//SoldTo:                        nil,
		// 付款人账单地址
		ShippingChargesPayment:        ShippingChargesPayment,
		//SmartPostDetail:               &rate.SmartPostShipmentDetail{
		//	Indicia:                    &_SmartPostIndiciaType,
		//	HubId:                      _hubId,
		//},
		PackageCount: &PkgCount,
		RequestedPackageLineItems:     GetPackages(),
	}

	_versionId := GetVersionId("crs",24)
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
		fmt.Printf("%+v\n",ret)
	}else{
		fmt.Printf("%+v\n",ret)
		fmt.Println("json: ")
		fmt.Println(string(byte))
	}
}