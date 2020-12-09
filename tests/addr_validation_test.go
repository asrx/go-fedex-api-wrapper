package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-fedex-api-wrapper/src/AddressValidationService"
	"github.com/asrx/go-fedex-api-wrapper/src/AddressValidationService/ComplexType"
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"testing"
	"time"
)

func Test_addrValidation(t *testing.T) {

	//var _url = "https://ws.fedex.com:443/web-services/addressvalidation"
	//var _url = "https://wsbeta.fedex.com:443/web-services/addressvalidation"

	_versionId := GetVersionId("aval", 4)

	addrList := []*ComplexType.AddressToValidate{
		&ComplexType.AddressToValidate{
			Address:           &ComplexType2.Address{
				StreetLines:           []string{"14022 S MADISON AVE"},
				City:                  "GLENPOOL",
				StateOrProvinceCode:   "OK",
				PostalCode:            "74033",
				CountryCode:           "US",
				Residential:           false,
			},
		},
	}

	addrRequest := ComplexType.AddressValidationRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 _versionId,
		InEffectAsOfTimestamp:   time.Time{},
		AddressesToValidate:     addrList,
	}

	c := AddressValidationService.NewAddressValidationPortType(false)
	ret, err := c.AddressValidation(&addrRequest)

	if err != nil {
		t.Errorf("Err： %v", err)
		return
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
