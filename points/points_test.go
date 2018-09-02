package main

import (
	"fmt"
	"testing"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

func TestLedger(t *testing.T) {
	// Various prints have been placed for testing purpose
	fmt.Println("Entered TestLedger Function")
	ledger := new(Ledger)
	stub := shim.NewMockStub("mockstub", ledger)
	if stub == nil {
		t.Fatalf("Mockstub creation failed")
	}

	//	var args1, args2, args3, args4, args5, args6 [][]byte
	var args1, args2, args3, args4, args5 [][]byte
	var response pb.Response
	var i int
	//====================
	//1 on
	//====================
	args_array1 := []string{"1", "13800138000", "on", "A", "", "1000", "", "", "", "2018-05-186T12:06:05Z"}
	//create the two dimensional byte array to pass to MockInvoke
	args1 = append(args1, []byte("writeUserTransaction")) //function name, pass it as the first byte array
	for i = 0; i < len(args_array1); i++ {
		args1 = append(args1, []byte(args_array1[i]))
	}
	fmt.Println("mock invoke args:", string(args1[1]))
	//Call the invoke function
	response = stub.MockInvoke("t1", args1)

	//Response handling
	printResponse("+++ on Response+++", response)

	//====================
	//2 exchange writeUserTransaction
	//====================
	args_array2 := []string{"2", "13800138000", "exchange", "A", "B", "100", "200", "1.0", "2.0", "2018-05-186T12:06:05Z"}
	//create the two dimensional byte array to pass to MockInvoke
	args2 = append(args2, []byte("writeUserTransaction")) //function name, pass it as the first byte array
	for i = 0; i < len(args_array2); i++ {
		args2 = append(args2, []byte(args_array2[i]))
	}
	fmt.Println("mock invoke args:", string(args2[1]))
	//Call the invoke function
	response = stub.MockInvoke("t2", args2)

	//Response handling
	printResponse("+++ exchange Response+++", response)
	//====================
	//3 off writeUserTransaction
	//====================
	args_array3 := []string{"3", "13800138000", "off", "B", "", "100", "", "", "", "2018-05-186T12:06:05Z"}
	//create the two dimensional byte array to pass to MockInvoke
	args3 = append(args3, []byte("writeUserTransaction")) //function name, pass it as the first byte array
	for i = 0; i < len(args_array3); i++ {
		args3 = append(args3, []byte(args_array3[i]))
	}
	fmt.Println("mock invoke args: ", string(args3[1]))
	//Call the invoke function
	response = stub.MockInvoke("t3", args3)

	//Response handling
	printResponse("+++ off Response+++", response)

	//====================
	//4 query getUserPointByMerchant
	//====================
	//args_array4 := []string{"13800138000", "A"}
	//create the two dimensional byte array to pass to MockInvoke
	//args4 = append(args4, []byte("getUserPointByMerchant")) //function name, pass it as the first byte array
	args_array4 := []string{"13800138000"}
	//create the two dimensional byte array to pass to MockInvoke
	args4 = append(args4, []byte("getUserPointAll")) //function name, pass it as the first byte array
	for i = 0; i < len(args_array4); i++ {
		args4 = append(args4, []byte(args_array4[i]))
	}
	//Call the invoke function
	response = stub.MockInvoke("t4", args4)

	//Response handling
	printResponse("+++ query Response+++", response)

	//====================
	//4 records
	//====================
	args_array5 := []string{"13800138000"}
	//create the two dimensional byte array to pass to MockInvoke
	args5 = append(args5, []byte("getUserTransactionAll")) //function name, pass it as the first byte array
	for i = 0; i < len(args_array5); i++ {
		args5 = append(args5, []byte(args_array5[i]))
	}
	//Call the invoke function
	response = stub.MockInvoke("t5", args5)

	//Response handling
	printResponse("+++ records Response+++", response)
}

func printResponse(heading string, response pb.Response) {

	fmt.Println(heading)

	fmt.Printf("Message: ")
	fmt.Println(response.GetMessage())
	fmt.Printf("Payload: ")
	fmt.Println(string(response.GetPayload()))
	fmt.Printf("Status: ")
	fmt.Println(response.GetStatus())

	fmt.Println("+++++End of Response+++++")
}
