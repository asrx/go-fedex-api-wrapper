package tests

import (
	ComplexType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/ComplexType"
	SimpleType2 "github.com/asrx/go-fedex-api-wrapper/src/Common/SimpleType"
)

// Groud
var ServiceType = SimpleType2.ServiceTypeFEDEX_GROUND

// HomeDelivery
var ServiceTypeHomeDelivery = SimpleType2.ServiceTypeGROUND_HOME_DELIVERY

// SmartPost
var ServiceTypeSmartPost = SimpleType2.ServiceTypeSMART_POST

// Priority_overnight
// FedEx将于下一个工作日上午10:30之前将货件递送到大部分美国地址；一些农村或偏远地区不晚于中午、下午4:30或5点；星期六递送不晚于中午、下午1:30或4:30。部分寄往阿拉斯加和夏威夷的货件将于两个工作日内送达。[服务枚举：FEDEX_PRIORITY_OVERNIGHT]
var ServiceTypePriority_overnight = SimpleType2.ServiceTypePRIORITY_OVERNIGHT

// Standard_overnight
// FedEx将于下一个工作日下午3点前将货件递送到大部分美国地址，部分乡间地区不晚于下午4:30（住宅地址不晚于晚8点）。[服务枚举：STANDARD_OVERNIGHT]
var ServiceTypeStandard_overnight = SimpleType2.ServiceTypeSTANDARD_OVERNIGHT

// FedEx将于第三个工作日下午4:30前将货件递送到大部分地区（住宅地址不晚于晚8点）。FedEx Express Saver不适用于电子邮件退件。[服务枚举：FEDEX_EXPRESS_SAVER]
var ServiceType_EXPRESS_SAVER = SimpleType2.ServiceTypeFEDEX_EXPRESS_SAVER

var PackageType = SimpleType2.PackagingTypeYOUR_PACKAGING

var DropoffType = SimpleType2.DropoffTypeREGULAR_PICKUP
var PreferredCurrency = "USD"
var PaymentType = SimpleType2.PaymentTypeSENDER
var ThirdPaymentType = SimpleType2.PaymentTypeTHIRD_PARTY

var _pkgCount = 1
var _GroupPackageCount uint = 1
var PkgCount uint = uint(_pkgCount)

var _weight float64 = 15

var _length uint = 11
var _width uint = 12
var _height uint = 13

// 签名
// var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT // 直接签名
var _signatureOptionType = SimpleType2.SignatureOptionTypeNO_SIGNATURE_REQUIRED // 不需要签名

var ShipFromContact = &ComplexType2.Contact{
	PersonName:   "ANL",
	CompanyName:  "AMERICAN NEW LOGISTICS",
	PhoneNumber:  "6262258083",
	EMailAddress: "OP1@LONGYUAN-LAX.COM",
}

var ShipFromAddr = &ComplexType2.Address{
	StreetLines:         []string{"244 Pine Barren Rd"},
	City:                "POOLER",
	StateOrProvinceCode: "GA",
	PostalCode:          "31322",
	// 只适用于波多黎各的地址。
	// UrbanizationCode:      "",
	CountryCode: "US",
	// CountryName:           "",
	Residential: false,
	// GeographicCoordinates: "",
}

// GA-POOLER
var ShipFromAddrGA = &ComplexType2.Address{
	StreetLines:         []string{"244 Pine Barren Rd"},
	City:                "POOLER",
	StateOrProvinceCode: "GA",
	PostalCode:          "31322",
	CountryCode:         "US",
	Residential:         false,
}

// TX-SOUTHLAKE
var ShipFromAddrTX = &ComplexType2.Address{
	StreetLines:         []string{"419 BANK ST", "STE 140"},
	City:                "SOUTHLAKE",
	StateOrProvinceCode: "TX",
	PostalCode:          "76092",
	CountryCode:         "US",
	Residential:         false,
}

// IN-PLAINFIELD
var ShipFromAddrIN = &ComplexType2.Address{
	StreetLines:         []string{"2575 E MAIN ST", "STE 100"},
	City:                "PLAINFIELD",
	StateOrProvinceCode: "IN",
	PostalCode:          "46168",
	CountryCode:         "US",
	Residential:         false,
}

// IL-
var ShipFromAddrIL = &ComplexType2.Address{
	StreetLines:         []string{"2475 TOUHY AVE"},
	City:                "ELK GROVE VILLAGE",
	StateOrProvinceCode: "IL",
	PostalCode:          "60007",
	CountryCode:         "US",
	Residential:         false,
}

// WA-SEATTLE
var ShipFromAddrWA = &ComplexType2.Address{
	StreetLines:         []string{"561 STRANDER BLVD"},
	City:                "SEATTLE",
	StateOrProvinceCode: "WA",
	PostalCode:          "98188",
	CountryCode:         "US",
	Residential:         false,
}

// CA2-HAYWARD
var ShipFromAddrCA2 = &ComplexType2.Address{
	StreetLines:         []string{"2802 W WINTON AVE"},
	City:                "HAYWARD",
	StateOrProvinceCode: "CA",
	PostalCode:          "94545",
	CountryCode:         "US",
	Residential:         false,
}

var ShipToContact = &ComplexType2.Contact{
	PersonName:   "ONT8",
	CompanyName:  "Amazon",
	PhoneNumber:  "8000000000",
	EMailAddress: "",
}

// 445 Valentine Industrial Pkwy, Pendergrass, GA 30567美国
var ShipToAddr = &ComplexType2.Address{
	StreetLines:         []string{"24300 Nandina Ave"},
	City:                "Moreno Valley",
	StateOrProvinceCode: "CA",
	PostalCode:          "92551",
	CountryCode:         "US",
	Residential:         false,
}

var Shipper = &ComplexType2.Party{
	AccountNumber: GetAccount(),
	Contact:       ShipFromContact,
	Address:       ShipFromAddrGA,
}

var Recipient = &ComplexType2.Party{
	Contact: ShipToContact,
	Address: ShipToAddr,
}

var ShippingChargesPayment = &ComplexType2.Payment{
	PaymentType: &PaymentType,
	Payor:       &ComplexType2.Payor{ResponsibleParty: Shipper},
}

// var _hubId = "5087"
// CA 生产
// var _hubId = "5902"
// NJ Test
var _hubId = "5087"

var _SmartPostIndiciaType = SimpleType2.SmartPostIndiciaTypePARCEL_SELECT
var SmartPostDetail = &ComplexType2.SmartPostShipmentDetail{
	Indicia: &_SmartPostIndiciaType,
	HubId:   _hubId,
}

var LabelFormatType_Common2D = SimpleType2.LabelFormatTypeCOMMON2D
var ImageType_PDF = SimpleType2.ShippingDocumentImageTypePDF
var LabelStockTypeSTOCK_4X6 = SimpleType2.LabelStockTypeSTOCK_4X6
var LabelSpecification = &ComplexType2.LabelSpecification{
	LabelFormatType: &LabelFormatType_Common2D,
	ImageType:       &ImageType_PDF,
	LabelStockType:  &LabelStockTypeSTOCK_4X6,
}

func GetPackages() []*ComplexType2.RequestedPackageLineItem {
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	var _customerReferenceValue = "#ref No 007"
	var _invType = SimpleType2.CustomerReferenceTypeINVOICE_NUMBER
	var _invValue = "#INV No 01"
	var _poType = SimpleType2.CustomerReferenceTypeP_O_NUMBER
	var _poValue = "#PO 01"
	var _deptType = SimpleType2.CustomerReferenceTypeDEPARTMENT_NUMBER
	var _deptValue = "#DEPT No 01"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{&_ptype}

	var _dimesions = &ComplexType2.Dimensions{
		Length: &_length,
		Width:  &_width,
		Height: &_height,
		Units:  &_dimensionsUnits,
	}

	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	for i := 1; i <= _pkgCount; i++ {
		var _SequenceNumber = uint(i)
		_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
			SequenceNumber:    &_SequenceNumber,
			GroupPackageCount: &_GroupPackageCount,
			// InsuredValue:                 nil,
			Weight: &ComplexType2.Weight{
				Units: &_WeightUnits,
				Value: _weight,
			},
			Dimensions: _dimesions,

			// CustomerReferences: &ComplexType2.CustomerReference{
			// 	CustomerReferenceType: &_customerReferenceType,
			// 	Value:                 _customerReferenceValue,
			// },

			CustomerReferences: []*ComplexType2.CustomerReference{
				{
					CustomerReferenceType: &_customerReferenceType,
					Value:                 _customerReferenceValue,
				},
				{
					CustomerReferenceType: &_invType,
					Value:                 _invValue,
				},
				{
					CustomerReferenceType: &_poType,
					Value:                 _poValue,
				},
				{
					CustomerReferenceType: &_deptType,
					Value:                 _deptValue,
				},
			},

			SpecialServicesRequested: &ComplexType2.PackageSpecialServicesRequested{
				SpecialServiceTypes: _specialServiceTypes,
				SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
					OptionType: &_signatureOptionType,
				},
			},
		})
	}

	return _packages
}

func GetPackagesForGroup() []*ComplexType2.RequestedPackageLineItem {
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	// var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	// var _customerReferenceValue = "xxxx"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{&_ptype}

	var _packages = []*ComplexType2.RequestedPackageLineItem{}

	var _SequenceNumber = uint(_pkgCount)
	_SequenceNumber = 1
	var _groupCount uint = 399

	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:    &_SequenceNumber,
		GroupPackageCount: &_groupCount,
		// InsuredValue:                 nil,
		Weight: &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: _weight,
		},
		Dimensions: &ComplexType2.Dimensions{
			Length: &_length,
			Width:  &_width,
			Height: &_height,
			Units:  &_dimensionsUnits,
		},
		// CustomerReferences: &ComplexType2.CustomerReference{
		// 	CustomerReferenceType: &_customerReferenceType,
		// 	Value:                 _customerReferenceValue,
		// },
		SpecialServicesRequested: &ComplexType2.PackageSpecialServicesRequested{
			SpecialServiceTypes: _specialServiceTypes,
			SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
				OptionType: &_signatureOptionType,
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
		SequenceNumber:    &_SequenceNumber2,
		GroupPackageCount: &_groupCount2,
		// InsuredValue:                 nil,
		Weight: &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: _weight2,
		},
		Dimensions: &ComplexType2.Dimensions{
			Length: &_length2,
			Width:  &_width2,
			Height: &_height2,
			Units:  &_dimensionsUnits,
		},
		// CustomerReferences: &ComplexType2.CustomerReference{
		// 	CustomerReferenceType: &_customerReferenceType,
		// 	Value:                 _customerReferenceValue,
		// },
		SpecialServicesRequested: &ComplexType2.PackageSpecialServicesRequested{
			SpecialServiceTypes: _specialServiceTypes,
			SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
				OptionType: &_signatureOptionType,
			},
		},
	})

	PkgCount = _groupCount + _groupCount2

	return _packages
}

func GetSubPackages() []*ComplexType2.RequestedPackageLineItem {
	var _SequenceNumber uint = 2
	var _GroupPackageCount uint = 1
	var _WeightUnits = SimpleType2.WeightUnitsLB

	var _dimensionsUnits = SimpleType2.LinearUnitsIN
	// var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	// var _customerReferenceValue = "yyy"

	var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION

	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{&_ptype}
	var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT

	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:    &_SequenceNumber,
		GroupPackageCount: &_GroupPackageCount,
		// InsuredValue:                 nil,
		Weight: &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: _weight,
		},
		Dimensions: &ComplexType2.Dimensions{
			Length: &_length,
			Width:  &_width,
			Height: &_height,
			Units:  &_dimensionsUnits,
		},
		// CustomerReferences: &ComplexType2.CustomerReference{
		// 	CustomerReferenceType: &_customerReferenceType,
		// 	Value:                 _customerReferenceValue,
		// },
		SpecialServicesRequested: &ComplexType2.PackageSpecialServicesRequested{
			SpecialServiceTypes: _specialServiceTypes,
			SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
				OptionType: &_signatureOptionType,
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
	// var _customerReferenceType = SimpleType2.CustomerReferenceTypeCUSTOMER_REFERENCE
	// var _customerReferenceValue = "yyyy"

	/*var _ptype = SimpleType2.PackageSpecialServiceTypeSIGNATURE_OPTION
	var _specialServiceTypes = []*SimpleType2.PackageSpecialServiceType{ &_ptype }
	var _signatureOptionType = SimpleType2.SignatureOptionTypeDIRECT*/

	var _packages = []*ComplexType2.RequestedPackageLineItem{}
	_packages = append(_packages, &ComplexType2.RequestedPackageLineItem{
		SequenceNumber:    &_SequenceNumber,
		GroupPackageCount: &_GroupPackageCount,
		// InsuredValue:                 nil,
		Weight: &ComplexType2.Weight{
			Units: &_WeightUnits,
			Value: 8,
		},
		Dimensions: &ComplexType2.Dimensions{
			Length: &_length,
			Width:  &_width,
			Height: &_height,
			Units:  &_dimensionsUnits,
		},
		// CustomerReferences: &ComplexType2.CustomerReference{
		// 	CustomerReferenceType: &_customerReferenceType,
		// 	Value:                 _customerReferenceValue,
		// },
		// SpecialServicesRequested:     &ComplexType2.PackageSpecialServicesRequested{
		//	SpecialServiceTypes:	_specialServiceTypes,
		//	SignatureOptionDetail: &ComplexType2.SignatureOptionDetail{
		//		OptionType:             &_signatureOptionType,
		//	},
		// },
	})

	return _packages
}
