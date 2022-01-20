package tests

import (
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

var ServiceType = SimpleType2.ServiceTypeFEDEX_GROUND
var ServiceTypeHomeDelivery = SimpleType2.ServiceTypeGROUND_HOME_DELIVERY
var ServiceTypeSmartPost = SimpleType2.ServiceTypeSMART_POST

var PackageType = SimpleType2.PackagingTypeYOUR_PACKAGING

var DropoffType = SimpleType2.DropoffTypeREGULAR_PICKUP
var PreferredCurrency = "USD"
var PaymentType = SimpleType2.PaymentTypeSENDER
var ThirdPaymentType = SimpleType2.PaymentTypeTHIRD_PARTY

var ShipFromContact = &ComplexType2.Contact{
	PersonName: "ANL",
	CompanyName: "AMERICAN NEW",
	PhoneNumber: "6262258083",
	EMailAddress: "OP1@LONGYUAN-LAX.COM",
}

var ShipFromAddr = &ComplexType2.Address{
	//StreetLines:           []string{"17539 HUGH LN"},
	//City:                  "LAND O LAKES",
	//StateOrProvinceCode:   "FL",
	//PostalCode:            "34638-7868",
	////只适用于波多黎各的地址。
	////UrbanizationCode:      "",
	//CountryCode:           "US",
	////CountryName:           "",
	//Residential:           true,
	////GeographicCoordinates: "",
	StreetLines:           []string{"45 Fernwood Ave Suite D"},
	City:                  "Edison",
	StateOrProvinceCode:   "NJ",
	PostalCode:            "08837",
	//只适用于波多黎各的地址。
	//UrbanizationCode:      "",
	CountryCode:           "US",
	//CountryName:           "",
	Residential:           false,
	//GeographicCoordinates: "",
}

var _pkgCount = 1
var _GroupPackageCount uint = 1
var PkgCount uint = uint(_pkgCount)

var _weight float64 = 56

var _length uint = 15
var _width uint = 18
var _height uint = 33

// 签名
//var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT // 直接签名
var _signatureOptionType = SimpleType2.SignatureOptionTypeNO_SIGNATURE_REQUIRED // 不需要签名

var ShipToContact = &ComplexType2.Contact{
	PersonName:          "VINCENT LU",
	CompanyName:         "ANL",
	PhoneNumber:         "8000000000",
	EMailAddress:        "VINCENT@AN-LOGISTICS.COM",
}
//var ShipToAddr = &ComplexType2.Address{
//	StreetLines:           []string{"12180 Hazelwood Dr"},
//	City:                  "Nokesville",
//	StateOrProvinceCode:   "VA",
//	PostalCode:            "20181",
//	CountryCode:           "US",
//	Residential:           false,
//}

// ONT8
//var ShipToAddr = &ComplexType2.Address{
//	StreetLines:           []string{"24300 Nandina Ave"},
//	City:                  "Moreno Valley",
//	StateOrProvinceCode:   "CA",
//	PostalCode:            "92551",
//	CountryCode:           "US",
//	Residential:           false,
//}

var ShipToAddr = &ComplexType2.Address{
	//StreetLines:           []string{"45 Fernwood Ave Suite D"},
	//City:                  "Edison",
	//StateOrProvinceCode:   "NJ",
	//PostalCode:            "08837",
	//CountryCode:           "US",
	//Residential:           false,
	StreetLines:           []string{"17539 HUGH LN"},
	City:                  "LAND O LAKES",
	StateOrProvinceCode:   "FL",
	PostalCode:            "34638-7868",
	CountryCode:           "US",
	Residential:           true,
}

var Shipper = &ComplexType2.Party{
	AccountNumber: GetAccount(),
	Contact:       ShipFromContact,
	Address:       ShipFromAddr,
}

var Recipient = &ComplexType2.Party{
	Contact:   ShipToContact,
	Address:   ShipToAddr,
}

var ShippingChargesPayment = &ComplexType2.Payment{
	PaymentType: &PaymentType,
	Payor:       &ComplexType2.Payor{ResponsibleParty: Shipper},
}

//var _hubId = "5087"
// CA 生产
//var _hubId = "5902"
// NJ Test
var _hubId = "5087"

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

func GetPackages() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "xxxx"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{ &_ptype }

	var _dimesions = &ComplexType2.Dimensions{
		Length: &_length,
		Width:  &_width,
		Height: &_height,
		Units:  &_dimensionsUnits,
	}

	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	for i := 1; i <= _pkgCount; i++ {
		_SequenceNumber = uint(_pkgCount)
		_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
			SequenceNumber:               &_SequenceNumber,
			GroupPackageCount:            &_GroupPackageCount,
			//InsuredValue:                 nil,
			Weight:                       &ComplexType2.Weight{
				Units: &_WeightUnits,
				Value: _weight,
			},
			Dimensions:                   _dimesions,
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
	}

	return _packages
}

func GetPackagesForGroup() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "xxxx"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{ &_ptype }

	var _packages = []*ComplexType2.RequestedPackageLineItem{}


	_SequenceNumber = uint(_pkgCount)
	_SequenceNumber = 1
	var _groupCount uint = 399

	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:               &_SequenceNumber,
		GroupPackageCount:            &_groupCount,
		//InsuredValue:                 nil,
		Weight:                       &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: _weight,
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

	var _SequenceNumber2 uint = 2
	var _groupCount2 uint = 600
	var _weight2 float64 = 5
	var _length2 uint = 5
	var _width2 uint = 5
	var _height2 uint = 5
	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:               &_SequenceNumber2,
		GroupPackageCount:            &_groupCount2,
		//InsuredValue:                 nil,
		Weight:                       &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: _weight2,
		},
		Dimensions:                   &ComplexType2.Dimensions{
			Length: &_length2,
			Width:  &_width2,
			Height: &_height2,
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

func GetSubPackages() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint = 2
	var _GroupPackageCount uint = 1
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "yyy"

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
			Value: _weight,
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

func GetPackagesSmartPost() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint = 1
	var _GroupPackageCount uint = 1
	var _WeightUnits = SimpleType2.WeightUnitsLB
	var _length uint = 10
	var _width uint = 7
	var _height uint = 8
	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "yyyy"

	/*var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION
	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{ &_ptype }
	var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT*/


	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:               &_SequenceNumber,
		GroupPackageCount:            &_GroupPackageCount,
		//InsuredValue:                 nil,
		Weight:                       &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: 8,
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
		//SpecialServicesRequested:     &ComplexType2.PackageSpecialServicesRequested{
		//	SpecialServiceTypes:	_specialServiceTypes,
		//	SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
		//		OptionType:             &_signatureOptionType,
		//	},
		//},
	})

	return _packages
}