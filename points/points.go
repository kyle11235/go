package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type Ledger struct {
}

type UserPoint struct {
	UserId     string `json:"uid"`
	MerchantId string `json:"merchantId"`
	Point      int32  `json:"point"`
}
type UserPointAll struct {
	Count int         `json:"count"`
	Items []UserPoint `json:"items"`
}
type PointExchangeRecord struct {
	TxId         string  `json:"transaction_id"`
	UserId       string  `json:"user_id"`
	Type         string  `json:"type"`
	FromMerchant string  `json:"from_merchant"`
	ToMerchant   string  `json:"to_merchant"`
	FromPoint    int32   `json:"from_point"`
	ToPoint      int32   `json:"to_point"`
	FromToken    float32 `json:"from_token"`
	ToToken      float32 `json:"to_token"`
	TxTime       string  `json:"createtime,omitempty"`
}
type TransactionAll struct {
	Count int                   `json:"count"`
	Items []PointExchangeRecord `json:"items"`
}

const (
	USER_INFO_PREFIX        = "user"
	USER_POINT_PREFIX       = "point"
	USER_TRASANCTION_PREFIX = "tx"

	TYPE_ON       = "on"
	TYPE_EXCHANGE = "exchange"
	TYPE_OFF      = "off"
)

type ExchangeRate struct {
	FromMerchant string
	ToMerchant   string
	Rate         string
}

const IndexExchangeRate = "From~To"

//==============
//Main Function
//==============
func main() {

	fmt.Println("Entered the Main Function")

	err := shim.Start(new(Ledger))
	if err != nil {
		fmt.Printf("Error starting  chaincode: %s", err)
	}
}

//=========================
//Initialize the chaincode
//=========================
func (t *Ledger) Init(stub shim.ChaincodeStubInterface) pb.Response {
	fmt.Println("Entering Init")
	return shim.Success(nil)
}

//========================
//Implementation of Invoke
//========================
func (t *Ledger) Invoke(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("Entering Invoke")
	function, args := stub.GetFunctionAndParameters()

	//List of Functions that can be invoked
	switch function {
	case "login":
		return t.login(stub, args)
	case "getUserPointByMerchant":
		return t.getUserPointByMerchant(stub, args)
	case "getUserPointAll":
		return t.getUserPointAll(stub, args)
	case "writeUserTransaction":
		return t.writeUserTransaction(stub, args)
	case "getUserTransactionAll":
		return t.getUserTransactionAll(stub, args)
	case "writeExchangeRate":
		return t.writeExchangeRate(stub, args)
	case "getExchangeRateByMerchant":
		return t.getExchangeRateByMerchant(stub, args)
	default:
		return shim.Error("Not a valid function " + function)
	}
}

func (t *Ledger) login(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Entering login")

	//Check if the number of arguments is 1
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments, expecting 2")
	}

	userid := args[0]
	passwd := args[1]
	//Create index based on
	compositeKey, err := stub.CreateCompositeKey(USER_INFO_PREFIX, []string{userid, passwd})
	if err != nil {
		return shim.Error(err.Error())
	}
	/*Insert the composite key into State DB
	value := []byte{0X00}
	err = stub.PutState(compositeKey, value)
	if err != nil {
		return shim.Error(err.Error())
	} */

	itemBytes, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	} else if itemBytes == nil {
		return shim.Error("Invalid userid/passwd: " + userid)
	}
	return shim.Success(itemBytes)
}

func (t *Ledger) getUserPointByMerchant(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	fmt.Println("Entering getUserPointByMerchant")

	//Check if the number of arguments is 1
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments, expecting 2")
	}

	userid := args[0]
	merchantid := args[1]
	//Create index based on userid merchant
	compositeKey, err := stub.CreateCompositeKey(USER_POINT_PREFIX, []string{userid, merchantid})
	if err != nil {
		return shim.Error(err.Error())
	}

	itemBytes, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error(err.Error())
	} else if itemBytes == nil {
		return shim.Error("Invalid userid/passwd: " + userid)
	}
	return shim.Success(itemBytes)
}

func (t *Ledger) getUserPointAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Entering getUserPointAll")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting user id to query")
	}
	pointIterator, err := stub.GetStateByPartialCompositeKey(USER_POINT_PREFIX, []string{args[0]})
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to getUserPointAll \"}"
		return shim.Error(jsonResp)
	}
	if pointIterator == nil {
		jsonResp := "{\"Error\":\"Nil  for getUserPointAll \"}"
		return shim.Error(jsonResp)
	}
	defer pointIterator.Close()

	var i int
	pointPayload := UserPointAll{0, make([]UserPoint, 0)}
	for i = 0; pointIterator.HasNext(); i++ {
		responseRange, err := pointIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// get the merchant name from composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		user := compositeKeyParts[0]
		merchant := compositeKeyParts[1]
		fmt.Printf("- found from index:%s  - %s-%s \n", objectType, user, merchant)
		fmt.Printf(" get value():%s\n", string(responseRange.GetValue()))
		pointRecord := UserPoint{}
		pointBytes := responseRange.GetValue()
		if pointBytes != nil {
			err = json.Unmarshal(pointBytes, &pointRecord)
		}
		pointPayload.Items = append(pointPayload.Items, pointRecord)
	}
	pointPayload.Count = i
	pointPayloadBytes, err := json.Marshal(pointPayload)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("getUserPointAll Response:%s\n", string(pointPayloadBytes))
	fmt.Println("end getUserPointAll")
	return shim.Success(pointPayloadBytes)
}

func updateUserPoint(stub shim.ChaincodeStubInterface, user string, merchant string, point int32) error {

	fmt.Println("Entering writeUserPoint")

	//Create index based on key1=userid key2=merchant
	compositeKey, err := stub.CreateCompositeKey(USER_POINT_PREFIX, []string{user, merchant})
	if err != nil {
		return err
	}
	userPointBytes, err := stub.GetState(compositeKey)
	if err != nil {
		fmt.Println("Failed to get " + compositeKey)
		return err
	}
	userPoint := UserPoint{}
	if userPointBytes != nil {
		err = json.Unmarshal(userPointBytes, &userPoint)
		if err != nil {
			return err
		}
		fmt.Printf("User old point value %s \n", string(userPointBytes))
	} else {
		userPoint.UserId = user
		userPoint.MerchantId = merchant
		userPoint.Point = 0
	}
	userPoint.Point += point
	userPointBytesNew, err := json.Marshal(&userPoint)
	fmt.Printf("User new point value %s \n", string(userPointBytesNew))
	err = stub.PutState(compositeKey, userPointBytesNew)
	if err != nil {
		return err
	}
	fmt.Printf("Put %s of %s point success.\n", user, merchant)
	return nil
}

func (t *Ledger) writeUserTransaction(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("-start writeUserTransaction")
	if len(args) != 10 {
		return shim.Error("Incorrect number of arguments. Expecting 10")
	}
	pointExchangeRecord := PointExchangeRecord{}
	pointExchangeRecord.TxId = args[0]
	pointExchangeRecord.UserId = args[1]
	pointExchangeRecord.Type = args[2]
	pointExchangeRecord.FromMerchant = args[3]
	pointExchangeRecord.ToMerchant = args[4]
	fromPoint, _ := strconv.ParseInt(args[5], 10, 32)
	pointExchangeRecord.FromPoint = int32(fromPoint)
	toPoint, _ := strconv.ParseInt(args[6], 10, 32)
	pointExchangeRecord.ToPoint = int32(toPoint)
	fromToken, _ := strconv.ParseFloat(args[7], 32)
	pointExchangeRecord.FromToken = float32(fromToken)
	toToken, _ := strconv.ParseFloat(args[8], 32)
	pointExchangeRecord.ToToken = float32(toToken)
	pointExchangeRecord.TxTime = args[9]

	fmt.Println("---step1 start")
	pointExchangeRecordBytes, err := json.Marshal(&pointExchangeRecord)
	fmt.Printf("%s \n", string(pointExchangeRecordBytes))
	txKey, err := stub.CreateCompositeKey(USER_TRASANCTION_PREFIX, []string{pointExchangeRecord.UserId, pointExchangeRecord.FromMerchant, pointExchangeRecord.TxId})
	//err = stub.PutState(pointExchangeRecord.TxId, pointExchangeRecordBytes)
	err = stub.PutState(txKey, pointExchangeRecordBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("---step1 end")

	fmt.Println("---step2 start")
	switch pointExchangeRecord.Type {
	case TYPE_ON:
		{
			fmt.Println("on")
			err = updateUserPoint(stub, pointExchangeRecord.UserId, pointExchangeRecord.FromMerchant, pointExchangeRecord.FromPoint)
		}
	case TYPE_EXCHANGE:
		{
			fmt.Println("out")
			err = updateUserPoint(stub, pointExchangeRecord.UserId, pointExchangeRecord.FromMerchant, -pointExchangeRecord.FromPoint)
			if err != nil {
				return shim.Error(err.Error())
			}
			fmt.Println("in")
			err = updateUserPoint(stub, pointExchangeRecord.UserId, pointExchangeRecord.ToMerchant, pointExchangeRecord.ToPoint)
		}
	case TYPE_OFF:
		{
			fmt.Println("off")
			err = updateUserPoint(stub, pointExchangeRecord.UserId, pointExchangeRecord.FromMerchant, -pointExchangeRecord.FromPoint)
		}
	}
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("---step2 end")

	fmt.Println("-end writeUserTransaction")
	return shim.Success(nil)
}

func (t *Ledger) getUserTransactionAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	fmt.Println("Entering getUserTransactionAll")
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting user id to query")
	}
	iterator, err := stub.GetStateByPartialCompositeKey(USER_TRASANCTION_PREFIX, []string{args[0]})
	if err != nil {
		jsonResp := "{\"Error\":\"Failed to getUserTransactionAll \"}"
		return shim.Error(jsonResp)
	}
	if iterator == nil {
		jsonResp := "{\"Error\":\"Nil  for getUserTransactionAll \"}"
		return shim.Error(jsonResp)
	}
	defer iterator.Close()

	var i int
	txPayload := TransactionAll{0, make([]PointExchangeRecord, 0)}
	for i = 0; iterator.HasNext(); i++ {
		responseRange, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		// get the merchant name from composite key
		// objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		// if err != nil {
		// 	return shim.Error(err.Error())
		// }
		// user := compositeKeyParts[0]
		// merchant := compositeKeyParts[1]
		// fmt.Printf("- found from index:%s  - %s-%s \n", objectType, user, merchant)
		fmt.Printf(" get value():%s\n", string(responseRange.GetValue()))
		txRecord := PointExchangeRecord{}
		txBytes := responseRange.GetValue()
		if txBytes != nil {
			err = json.Unmarshal(txBytes, &txRecord)
		}
		txPayload.Items = append(txPayload.Items, txRecord)
	}
	txPayload.Count = i
	txPayloadBytes, err := json.Marshal(txPayload)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Printf("getUserTransactionAll Response:%s\n", string(txPayloadBytes))
	fmt.Println("end getUserTransactionAll")
	return shim.Success(txPayloadBytes)
}

func (t *Ledger) writeExchangeRate(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}

	fromMerchant := args[0]
	toMerchant := args[1]
	rate := args[2]

	exchangeRateIndexKey, err := stub.CreateCompositeKey(IndexExchangeRate, []string{fromMerchant, toMerchant})
	if err != nil {
		return shim.Error(err.Error())
	}

	stub.PutState(exchangeRateIndexKey, []byte(rate))

	return shim.Success(nil)
}

func (t *Ledger) getExchangeRateByMerchant(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	fromMerchant := args[0]
	exchangeRateResultsIterator, err := stub.GetStateByPartialCompositeKey(IndexExchangeRate, []string{fromMerchant})

	if err != nil {
		return shim.Error(err.Error())
	}

	defer exchangeRateResultsIterator.Close()

	// buffer is a JSON array containing QueryResults
	var buffer bytes.Buffer
	buffer.WriteString("\"items\":[")

	bArrayMemberAlreadyWritten := false

	// Iterate through result set and for each FromMerchant found
	var i int
	for i = 0; exchangeRateResultsIterator.HasNext(); i++ {
		responseRange, err := exchangeRateResultsIterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}

		// get ToMerchant from composite key
		objectType, compositeKeyParts, err := stub.SplitCompositeKey(responseRange.Key)
		if err != nil {
			return shim.Error(err.Error())
		}
		if objectType != IndexExchangeRate {
			return shim.Error(err.Error())
		}
		//returnedFromMerchant := compositeKeyParts[0]
		returnedToMerchant := compositeKeyParts[1]

		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"to\":")
		buffer.WriteString("\"")
		buffer.WriteString(returnedToMerchant)
		buffer.WriteString("\"")

		buffer.WriteString(", \"rate\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(responseRange.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getExchangeRateByMerchant queryResult:\n%s\n", buffer.String())

	return shim.Success(buffer.Bytes())
}
