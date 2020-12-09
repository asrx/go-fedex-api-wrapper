package tests

import (
	calc "github.com/asrx/go-fedex-api-wrapper/src/Demo"
	"github.com/hooklift/gowsdl/soap"
	"log"
	"testing"
)

func Test_Demo(t *testing.T)  {
	log.Println("测试...测试...测试...")
}

func Test_calc(t *testing.T) {
	client := soap.NewClient("http://www.dneonline.com/calculator.asmx")
	c := calc.NewCalculatorSoap(client)
	r, err := c.Add(&calc.Add{
		IntA:    2,
		IntB:    3,
	})

	if err != nil {
		t.Error(err)
		return
	}

	t.Logf("Result: %d\n", r.AddResult)
}