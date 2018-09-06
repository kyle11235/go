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
// createCoupon
// ===========================================================
func TestCreateCoupon(t *testing.T) {
	// method
	method := "createCoupon"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"couponID", "couponName", "999", "100"}

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
// moveCouponToUser
// ===========================================================
func TestMoveCouponToUser(t *testing.T) {
	// method
	method := "moveCouponToUser"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"couponID", "10", "userID"}

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
// getCoupon
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
// createOrder
// ===========================================================
func TestCreateOrder(t *testing.T) {
	// method
	method := "createOrder"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"orderID", "userID", "couponID", "5"}

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
// auditOrder
// ===========================================================
func TestAuditOrder(t *testing.T) {
	// method
	method := "auditOrder"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"orderID"}

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
// getOrder
// ===========================================================
func TestGetOrder(t *testing.T) {
	// method
	method := "getOrder"
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
// getUserCoupon
// ===========================================================
func TestGetUserCoupon(t *testing.T) {
	// method
	method := "getUserCoupon"
	fmt.Println("------ start ------ ", method)

	// other args
	argsArray := []string{"userID"}

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

// print
func printResponse(response pb.Response) {
	fmt.Printf("- status=")
	fmt.Println(response.GetStatus())
	fmt.Printf("- error message=")
	fmt.Println(response.GetMessage())
	fmt.Printf("- payload=")
	fmt.Println(string(response.GetPayload()))
}

// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"createCoupon","args":["couponID", "couponName", "999", "100"],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"moveCouponToUser","args":["couponID", "10", "userID"],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getCoupon","args":[],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getUserCoupon","args":["userID"],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"createOrder","args":["orderID", "userID", "couponID", "5"],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"auditOrder","args":["orderID"],"chaincodeVer":"v1"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getOrder","args":[],"chaincodeVer":"v1"}'
