package tests

import (
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	"github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
	"github.com/asrx/go-fedex-api-wrapper/src/CountryService"
	"github.com/asrx/go-fedex-api-wrapper/src/CountryService/ComplexType"
	"encoding/json"
	"fmt"
	"testing"
	"time"
)

func Test_Country(t *testing.T)  {
	_carrierCode := SimpleType.CarrierCodeTypeFDXE
	_ValidatePostalRequest := ComplexType.ValidatePostalRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		Version:                 GetVersionId("cnty", 8),
		ShipDateTime:			time.Now(),
		Address:           		&ComplexType2.Address{
			PostalCode:            "74033",
			CountryCode:           "US",
		},
		CarrierCode: 			 &_carrierCode,
	}

	c := CountryService.NewCountryPortType(false)
	ret, err := c.ValidatePostal(&_ValidatePostalRequest)

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

	fmt.Println(_ValidatePostalRequest)
	fmt.Println("country service")
}
