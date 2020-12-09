package tests

import (
	"encoding/json"
	"fmt"
	"github.com/asrx/go-fedex-api-wrapper/src/AsyncService"
	"github.com/asrx/go-fedex-api-wrapper/src/AsyncService/ComplexType"
	"testing"
)

// Not through
func Test_Async(t *testing.T) {
	_jobId := "794660513970"

	request := &ComplexType.RetrieveJobResultsRequest{
		WebAuthenticationDetail: GetWebAuthenticationDetail(),
		ClientDetail:            GetClientDetail(),
		//TransactionDetail:       nil,
		Version:                 GetVersionId("async",4),
		JobId:                   _jobId,
		//Filters:                 nil,
	}

	c := AsyncService.NewASYNCPortType(true)
	ret, err := c.RetrieveJobResults(request)

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
