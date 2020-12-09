package tests

import "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"

const (
	_KEY = ""
	_PWD = ""
	_ACCOUNT = ""
	_METER = ""

)

func GetAccount() string {
	return _ACCOUNT
}
func GetClientDetail() *ComplexType.ClientDetail {
	return &ComplexType.ClientDetail{
		AccountNumber: _ACCOUNT,
		MeterNumber:   _METER,
	}
}

func GetWebAuthenticationDetail() *ComplexType.WebAuthenticationDetail {
	return &ComplexType.WebAuthenticationDetail{
		UserCredential:   &ComplexType.WebAuthenticationCredential{
			Key:      _KEY,
			Password: _PWD,
		},
	}
}

var I32 int32 = 0
func GetVersionId(serviceId string, major int32) *ComplexType.VersionId {
	return &ComplexType.VersionId{
		ServiceId:    serviceId,
		Major:        major,
		Intermediate: &I32,
		Minor:        &I32,
	}
}