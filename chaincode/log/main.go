package main

import (
	"bytes"
	"crypto/md5"
	"crypto/rand"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// log
type Log struct {
	DocType      string        `json:"docType"`
	ID           string        `json:"id"` // email
	Name         string        `json:"name"`
	Desc         string        `json:"desc"`
	YearInducted string        `json:"yearInducted"`
	Photo        string        `json:"photo"`
	RunnerAwards []RunnerAward `json:"runnerAwards"`
}

type RunnerAward struct {
	DocType      string `json:"docType"`
	ID           string `json:"id"` // guid
	LogID        string `json:"logID"`
	Title        string `json:"title"`
	MarathonName string `json:"marathonName"`
	Type         string `json:"type"`
	Time         string `json:"time"`
}

// log list
type LogList struct {
	Items []Log `json:"items"`
}

// prefix
const (
	PrefixLog         = "log"
	PrefixRunnerAward = "runnerAward"
)

// main
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// init
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	return shim.Success(nil)
}

// invoke
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	fmt.Println("invoking " + function)

	// functions
	if function == "putLog" {
		return t.putLog(stub, args)
	} else if function == "putRunnerAward" {
		return t.putRunnerAward(stub, args)
	} else if function == "getLogList" {
		return t.getLogList(stub, args)
	}

	// result
	message := "invoke did not find func: " + function
	fmt.Println(message)
	return shim.Error(message)
}

// ===========================================================
// putLog
// ===========================================================
func (t *SimpleChaincode) putLog(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// check args
	count := 5
	if len(args) != count {
		return shim.Error(fmt.Sprintf("Incorrect number of arguments. Expecting %d", count))
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(fmt.Sprintf("argument must be a non-empty string, index=%d", i))
		}
	}

	id := args[0]
	name := args[1]
	desc := args[2]
	yearInducted := args[3]
	photo := args[4]

	// create
	key, err := stub.CreateCompositeKey(PrefixLog, []string{id})
	if err != nil {
		return shim.Error(err.Error())
	}
	log := &Log{PrefixLog, id, name, desc, yearInducted, photo, nil}
	bytes, err := json.Marshal(log)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("create log success")
	return shim.Success(nil)
}

// ===========================================================
// putRunnerAward
// ===========================================================
func (t *SimpleChaincode) putRunnerAward(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// check args
	count := 5
	if len(args) != count {
		return shim.Error(fmt.Sprintf("Incorrect number of arguments. Expecting %d", count))
	}

	for i := 0; i < len(args); i++ {
		if len(args[i]) <= 0 {
			return shim.Error(fmt.Sprintf("argument must be a non-empty string, index=%d", i))
		}
	}

	id := getGUID()
	logID := args[0]
	title := args[1]
	marathonName := args[2]
	marathonType := args[3]
	time := args[4]

	// create
	key, err := stub.CreateCompositeKey(PrefixRunnerAward, []string{id})
	if err != nil {
		return shim.Error(err.Error())
	}
	runnerAward := &RunnerAward{PrefixRunnerAward, id, logID, title, marathonName, marathonType, time}
	bytes, err := json.Marshal(runnerAward)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("create runnerAward success")
	return shim.Success(nil)
}

// ===========================================================
// getLogList
// ===========================================================
func (t *SimpleChaincode) getLogList(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	iterator, err := stub.GetStateByPartialCompositeKey(PrefixLog, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	if iterator == nil {
		return shim.Error("getLogList error")
	}
	defer iterator.Close()

	list := LogList{make([]Log, 0)}
	for iterator.HasNext() {
		next, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		item := Log{}
		itemBytes := next.GetValue()
		if itemBytes != nil {
			err = json.Unmarshal(itemBytes, &item)
		}

		// add award list
		item.RunnerAwards, err = getRunnerAwardList(stub, item.ID)
		if err != nil {
			return shim.Error(err.Error())
		}

		list.Items = append(list.Items, item)
	}
	listBytes, err := json.Marshal(list)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("get log success")
	return shim.Success(listBytes)
}

// ===========================================================
// get runner award list
// ===========================================================
func getRunnerAwardList(stub shim.ChaincodeStubInterface, logID string) ([]RunnerAward, error) {
	queryString := "{\"selector\":{\"logID\":\"" + logID + "\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return nil, err
	}

	runnerAwardList := make([]RunnerAward, 0)
	err = json.Unmarshal(queryResults, &runnerAwardList)
	if err != nil {
		return nil, err
	}

	return runnerAwardList, nil
}

// =========================================================================================
// private  - rich query - getQueryResultForQueryString executes the passed in query string.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

// md5
func getMd5(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

// guid
func getGUID() string {
	b := make([]byte, 48)

	if _, err := io.ReadFull(rand.Reader, b); err != nil {
		return ""
	}
	return getMd5(base64.URLEncoding.EncodeToString(b))
}
