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
	call("createCoupon", []string{"couponID", "couponName", "999", "100"})
	call("moveCouponToUser", []string{"couponID", "10", "userID"})
	call("getCoupon", nil)
	call("createOrder", []string{"orderID", "userID", "couponID", "5"})
	call("auditOrder", []string{"orderID"})
	call("getOrder", nil)
	call("getUserCoupon", []string{"userID"})
	call("deleteAll", nil)
	call("getCoupon", nil)
}

// ===========================================================
// private call
// ===========================================================
func call(method string, argsArray []string) {
	fmt.Println("------ start ------ ", method)
	// all args
	var args [][]byte
	args = append(args, []byte(method))
	if argsArray != nil {
		if len(argsArray) > 0 {
			fmt.Printf("-")
		}
		for i := 0; i < len(argsArray); i++ {
			args = append(args, []byte(argsArray[i]))
			fmt.Printf("p%d=%s, ", i, args[i])
		}
	}
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

// curl
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"createCoupon","args":["couponID", "couponName", "999", "100"],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"moveCouponToUser","args":["couponID", "10", "userID"],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getCoupon","args":[],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getUserCoupon","args":["userID"],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"createOrder","args":["orderID", "userID", "couponID", "5"],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/invocation -d '{"channel":"test1orderer","chaincode":"points1","method":"auditOrder","args":["orderID"],"chaincodeVer":"v2"}'
// curl -H "Content-type:application/json" -X POST http://129.213.123.198:4111/bcsgw/rest/v1/transaction/query -d '{"channel":"test1orderer","chaincode":"points1","method":"getOrder","args":[],"chaincodeVer":"v2"}'

// cli
// peer chaincode invoke -n mycc -c '{"Args":["createCoupon", "couponID", "couponName", "999", "100"]}' -C myc
// peer chaincode invoke -n mycc -c '{"Args":["getCoupon"]}' -C myc
