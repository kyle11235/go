package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"

	"testing"
)

var simpleAsset = new(SimpleChaincode)
var stub = shim.NewMockStub("mockstub", simpleAsset)

// ===========================================================
// test
// ===========================================================
func Test(t *testing.T) {
	call("putLog", []string{"x.y.z@example.com", "kyle zhang", "desc", "2019", "https://upload.wikimedia.org/wikipedia/commons/9/9a/Joan_Benoit_2008.jpg"})
	call("putRunnerAward", []string{"x.y.z@example.com", "Top 10 in Melbourne marathon 2019", "2019 Melbourne Marathon Festival", "half marathon", "58:05"})
	call("putRunnerAward", []string{"x.y.z@example.com", "Top 40 in Tokyo marathon 2018", "IOC Tokyo Olympic marathon", "full marathon", "2:22:01"})
	call("getLogList", nil)
}

// ===========================================================
// private call
// ===========================================================
func call(method string, argsArray []string) {
	fmt.Println("------ start ------ ", method)
	// all args
	var args [][]byte
	args = append(args, []byte(method))
	fmt.Printf("- args=[")
	fmt.Printf("p0=%s", method)
	if argsArray != nil {
		for i := 0; i < len(argsArray); i++ {
			args = append(args, []byte(argsArray[i]))
			fmt.Printf(",p%d=%s", i+1, argsArray[i])
		}
	}
	fmt.Printf("]")
	fmt.Println("")
	// invoke
	response := stub.MockInvoke("uuid", args)
	fmt.Printf("- status=")
	fmt.Println(response.GetStatus())
	fmt.Printf("- error message=")
	fmt.Println(response.GetMessage())
	fmt.Printf("- payload=")
	fmt.Println(string(response.GetPayload()))
	fmt.Println("------ end ------ ")
	fmt.Println("")
}
