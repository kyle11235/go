package main

import (
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"

	"testing"
)

var simpleAsset = new(SimpleChaincode)
var stub = shim.NewMockStub("mockstub", simpleAsset)

// ===========================================================
func TestCreateCoupon(t *testing.T) {
	// method
	method := "createCoupon"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"101", "300-50", "999", "100"}

	// all args
	var args [][]byte
	args = append(args, []byte(method))
	if len(argsArray) > 0 {
		fmt.Printf("-")
	}
	for i := 0; i < len(argsArray); i++ {
		args = append(args, []byte(argsArray[i]))
		fmt.Printf("p%d=%s, ", i, args[i])
	}

	// invoke
	response := stub.MockInvoke("uuid", args)
	printResponse(response)

	fmt.Println("------  end  ------")
}

// ===========================================================
func TestMoveCouponToUser(t *testing.T) {
	// method
	method := "moveCouponToUser"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"101", "10", "kyle"}

	// all args
	var args [][]byte
	args = append(args, []byte(method))
	if len(argsArray) > 0 {
		fmt.Printf("-")
	}
	for i := 0; i < len(argsArray); i++ {
		args = append(args, []byte(argsArray[i]))
		fmt.Printf("p%d=%s, ", i, args[i])
	}

	// invoke
	response := stub.MockInvoke("uuid", args)
	printResponse(response)

	fmt.Println("------  end  ------")
}

// ===========================================================
func TestGetCoupon(t *testing.T) {
	// method
	method := "getCoupon"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{}

	// all args
	var args [][]byte
	args = append(args, []byte(method))
	if len(argsArray) > 0 {
		fmt.Printf("-")
	}
	for i := 0; i < len(argsArray); i++ {
		args = append(args, []byte(argsArray[i]))
		fmt.Printf("p%d=%s, ", i, args[i])
	}

	// invoke
	response := stub.MockInvoke("uuid", args)
	printResponse(response)

	fmt.Println("------  end  ------")
}

// ===========================================================
func TestGetUserCoupon(t *testing.T) {
	// method
	method := "getUserCoupon"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"kyle"}

	// all args
	var args [][]byte
	args = append(args, []byte(method))
	if len(argsArray) > 0 {
		fmt.Printf("-")
	}
	for i := 0; i < len(argsArray); i++ {
		args = append(args, []byte(argsArray[i]))
		fmt.Printf("p%d=%s, ", i, args[i])
	}

	// invoke
	response := stub.MockInvoke("uuid", args)
	printResponse(response)

	fmt.Println("------  end  ------")
}

// ===========================================================
func printResponse(response pb.Response) {
	fmt.Printf("- status=")
	fmt.Println(response.GetStatus())
	fmt.Printf("- error message=")
	fmt.Println(response.GetMessage())
	fmt.Printf("- payload=")
	fmt.Println(string(response.GetPayload()))
}
