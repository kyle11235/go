package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"testing"
)

var simpleAsset = new(SimpleAsset)
var stub = shim.NewMockStub("mockstub", simpleAsset)

func TestSet(t *testing.T) {
	method := "set"
	fmt.Println("--- start testing... ", method)

	var args [][]byte
	var response pb.Response
	var i int

	// args
	args_array := []string{"A", "kyle"}

	// first arg is method name
	args = append(args, []byte(method))
	for i = 0; i < len(args_array); i++ {
		args = append(args, []byte(args_array[i]))
	}
	fmt.Println("args=", string(args[1]))

	response = stub.MockInvoke("t1", args)
	printResponse("response=", response)

	fmt.Println("--- end testing")
}

func TestGet(t *testing.T) {
	method := "get"
	fmt.Println("--- start testing... ", method)

	var args [][]byte
	var response pb.Response
	var i int

	// args
	args_array := []string{"A"}

	// first arg is method name
	args = append(args, []byte(method))
	for i = 0; i < len(args_array); i++ {
		args = append(args, []byte(args_array[i]))
	}
	fmt.Println("args=", string(args[1]))

	response = stub.MockInvoke("t1", args)
	printResponse("response=", response)

	fmt.Println("--- end testing")
}

func printResponse(heading string, response pb.Response) {

	fmt.Println(heading)

	fmt.Printf("message=")
	fmt.Println(response.GetMessage())
	fmt.Printf("payload=")
	fmt.Println(string(response.GetPayload()))
	fmt.Printf("status=")
	fmt.Println(response.GetStatus())

}
