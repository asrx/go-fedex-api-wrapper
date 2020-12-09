package tests

import (
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

var ServiceType = SimpleType2.ServiceTypeFEDEX_GROUND
var PackageType = SimpleType2.PackagingTypeYOUR_PACKAGING

var DropoffType = SimpleType2.DropoffTypeREGULAR_PICKUP
var PreferredCurrency = "USD"
var PaymentType = SimpleType2.PaymentTypeSENDER


var Shipper = &ComplexType2.Party{
	AccountNumber: GetAccount(),
	Contact:       &ComplexType2.Contact{
	PersonName:          "Donovan",
	CompanyName:         "ANL",
	PhoneNumber:         "6262258083",
	EMailAddress:        "xudong0226@163.com",
	},
	Address:       &ComplexType2.Address{
	StreetLines:           []string{"16018 Adelante st Suite D"},
	City:                  "Irwindale",
	StateOrProvinceCode:   "CA",
	PostalCode:            "91702",
	//只适用于波多黎各的地址。
	//UrbanizationCode:      "",
	CountryCode:           "US",
	//CountryName:           "",
	Residential:           false,
	//GeographicCoordinates: "",
	},
}

var Recipient = &ComplexType2.Party{
	Contact:       &ComplexType2.Contact{
		PersonName:          "Alex",
		CompanyName:         "ANL",
		PhoneNumber:         "8000000000",
		EMailAddress:        "xudong0226@163.com",
	},
	Address:       &ComplexType2.Address{
		StreetLines:           []string{"401 Independence Rd"},
		City:                  "FLORENCE",
		StateOrProvinceCode:   "NJ",
		PostalCode:            "08518",
		CountryCode:           "US",
		Residential:           false,
	},
}

var ShippingChargesPayment = &ComplexType2.Payment{
	PaymentType: &PaymentType,
	Payor:       &ComplexType2.Payor{ResponsibleParty: Shipper},
}

var _hubId = "5902"
var _SmartPostIndiciaType = SimpleType2.SmartPostIndiciaTypePARCEL_SELECT
var SmartPostDetail = &ComplexType2.SmartPostShipmentDetail{
	Indicia:                    &_SmartPostIndiciaType,
	HubId:                      _hubId,
}

var LabelFormatType_Common2D = SimpleType2.LabelFormatTypeCOMMON2D
var ImageType_PDF = SimpleType2.ShippingDocumentImageTypePDF
var LabelStockTypeSTOCK_4X6 = SimpleType2.LabelStockTypeSTOCK_4X6
var LabelSpecification = &ComplexType2.LabelSpecification{
	LabelFormatType:          &LabelFormatType_Common2D,
	ImageType:                &ImageType_PDF,
	LabelStockType:           &LabelStockTypeSTOCK_4X6,
}
var PkgCount uint = 1

func GetPackages() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint = 1
	var _GroupPackageCount uint = 1
	var _WeightUnits = SimpleType2.WeightUnitsLB
	var _length uint = 40
	var _width uint = 30
	var _height uint = 20
	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "xxxx"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{ &_ptype }
	var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT


	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:               &_SequenceNumber,
		GroupPackageCount:            &_GroupPackageCount,
		//InsuredValue:                 nil,
		Weight:                       &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: 20,
		},
		Dimensions:                   &ComplexType2.Dimensions{
			Length: &_length,
			Width:  &_width,
			Height: &_height,
			Units:  &_dimensionsUnits,
		},
		CustomerReferences:           &ComplexType2.CustomerReference{
			CustomerReferenceType: &_customerReferenceType,
			Value:                 _customerReferenceValue,
		},
		SpecialServicesRequested:     &ComplexType2.PackageSpecialServicesRequested{
			SpecialServiceTypes:	_specialServiceTypes,
			SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
				OptionType:             &_signatureOptionType,
			},
		},
	})

	return _packages
}