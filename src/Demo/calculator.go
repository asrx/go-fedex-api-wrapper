package Demo

import (
	"encoding/xml"
	"github.com/hooklift/gowsdl/soap"
	"time"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

type Add struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Add"`

	IntA int32 `xml:"intA,omitempty"`

	IntB int32 `xml:"intB,omitempty"`
}

type AddResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ AddResponse"`

	AddResult int32 `xml:"AddResult,omitempty"`
}

type Subtract struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Subtract"`

	IntA int32 `xml:"intA,omitempty"`

	IntB int32 `xml:"intB,omitempty"`
}

type SubtractResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ SubtractResponse"`

	SubtractResult int32 `xml:"SubtractResult,omitempty"`
}

type Multiply struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Multiply"`

	IntA int32 `xml:"intA,omitempty"`

	IntB int32 `xml:"intB,omitempty"`
}

type MultiplyResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ MultiplyResponse"`

	MultiplyResult int32 `xml:"MultiplyResult,omitempty"`
}

type Divide struct {
	XMLName xml.Name `xml:"http://tempuri.org/ Divide"`

	IntA int32 `xml:"intA,omitempty"`

	IntB int32 `xml:"intB,omitempty"`
}

type DivideResponse struct {
	XMLName xml.Name `xml:"http://tempuri.org/ DivideResponse"`

	DivideResult int32 `xml:"DivideResult,omitempty"`
}

type CalculatorSoap interface {

	/* Adds two integers. This is a test WebService. ©DNE Online */
	Add(request *Add) (*AddResponse, error)

	Subtract(request *Subtract) (*SubtractResponse, error)

	Multiply(request *Multiply) (*MultiplyResponse, error)

	Divide(request *Divide) (*DivideResponse, error)
}

type calculatorSoap struct {
	client *soap.Client
}

func NewCalculatorSoap(client *soap.Client) CalculatorSoap {
	return &calculatorSoap{
		client: client,
	}
}

func (service *calculatorSoap) Add(request *Add) (*AddResponse, error) {
	response := new(AddResponse)
	err := service.client.Call("http://tempuri.org/Add", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *calculatorSoap) Subtract(request *Subtract) (*SubtractResponse, error) {
	response := new(SubtractResponse)
	err := service.client.Call("http://tempuri.org/Subtract", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *calculatorSoap) Multiply(request *Multiply) (*MultiplyResponse, error) {
	response := new(MultiplyResponse)
	err := service.client.Call("http://tempuri.org/Multiply", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *calculatorSoap) Divide(request *Divide) (*DivideResponse, error) {
	response := new(DivideResponse)
	err := service.client.Call("http://tempuri.org/Divide", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}
