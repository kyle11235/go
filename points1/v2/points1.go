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
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// coupon
type Coupon struct {
	ObjectType string `json:"docType"`
	ID         string `json:"id"`
	Name       string `json:"name"`
	Points     int    `json:"points"`
	Total      int    `json:"total"`
	Used       int    `json:"used"`
}

// coupon list
type CouponList struct {
	Items []Coupon `json:"items"`
}

// user coupon
type UserCoupon struct {
	UserID   string `json:"userID"`
	CouponID string `json:"couponID"`
	Name     string `json:"couponName"`
	Points   int    `json:"points"`
	Total    int    `json:"total"`
	Used     int    `json:"used"`
}

// user coupon list
type UserCouponList struct {
	Items []UserCoupon `json:"items"`
}

// order
type Order struct {
	ObjectType  string `json:"docType"`
	ID          string `json:"id"`
	UserID      string `json:"userID"`
	CouponName  string `json:"couponName"`
	CouponCount int    `json:"couponCount"`
	Status      string `json:"status"` // locked/audited
}

// order list
type OrderList struct {
	Items []Order `json:"items"`
}

// prefix
const (
	PREFIX_COUPON      = "coupon2"
	PREFIX_USER_COUPON = "user-coupon"
	PREFIX_ORDER       = "order"
	PREFIX_USER_ORDER  = "user-order"
	STATUS_LOCKED      = "locked"
	STATUS_AUDITED     = "audited"
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
	if function == "createCoupon" {
		return t.createCoupon(stub, args)
	} else if function == "getCoupon" {
		return t.getCoupon(stub, args)
	} else if function == "moveCouponToUser" {
		return t.moveCouponToUser(stub, args)
	} else if function == "getUserCoupon" {
		return t.getUserCoupon(stub, args)
	} else if function == "createOrder" {
		return t.createOrder(stub, args)
	} else if function == "getOrder" {
		return t.getOrder(stub, args)
	} else if function == "auditOrder" {
		return t.auditOrder(stub, args)
	} else if function == "deleteAll" {
		return t.deleteAll(stub, args)
	}

	// result
	message := "invoke did not find func: " + function
	fmt.Println(message)
	return shim.Error(message)
}

// ===========================================================
// createCoupon
// ===========================================================
func (t *SimpleChaincode) createCoupon(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// check args
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty number")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty number")
	}

	id := args[0]
	name := args[1]
	points, err := strconv.Atoi(args[2])
	if err != nil {
		return shim.Error("3rd argument must be a number")
	}
	count, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("4th argument must be a number")
	}

	// check existence
	var coupon *Coupon
	coupon, err = getCouponByID(stub, id)
	if err != nil {
		return shim.Error("Failed to get coupon: " + err.Error())
	}
	if coupon != nil {
		fmt.Println("This coupon already exists: " + id)
		return shim.Error("This coupon already exists: " + id)
	}

	// create
	key, err := stub.CreateCompositeKey(PREFIX_COUPON, []string{id})
	if err != nil {
		return shim.Error(err.Error())
	}
	coupon = &Coupon{PREFIX_COUPON, id, name, points, count, 0}
	bytes, err := json.Marshal(coupon)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(key, bytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("create coupon success")
	return shim.Success(nil)
}

// ===========================================================
// getCoupon
// ===========================================================
func (t *SimpleChaincode) getCoupon(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	queryString := "{\"selector\":{\"docType\":\"" + PREFIX_COUPON + "\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===========================================================
// moveCouponToUser
// ===========================================================
func (t *SimpleChaincode) moveCouponToUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	// check args
	if len(args) != 3 {
		return shim.Error("Incorrect number of arguments. Expecting 3")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty number")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}

	couponID := args[0]
	count, err := strconv.Atoi(args[1])
	if err != nil {
		return shim.Error("2rd argument must be a number")
	}
	userID := args[2]

	// get coupon
	var coupon *Coupon
	coupon, err = getCouponByID(stub, couponID)
	if err != nil {
		return shim.Error("Failed to get coupon: " + err.Error())
	}
	if coupon == nil {
		fmt.Println("This coupon NOT exists: " + couponID)
		return shim.Error("This coupon NOT exists: " + couponID)
	}

	// update coupon
	coupon.Used = coupon.Used + count
	err = updateCoupon(stub, coupon)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get user coupon - call a private function
	var userCoupon *UserCoupon
	userCoupon, err = getUserCouponByID(stub, userID, couponID)

	if err != nil {
		return shim.Error("Failed to get user coupon: " + err.Error())
	}
	if userCoupon != nil {
		//update
		userCoupon.Total = userCoupon.Total + count
	} else {
		// create
		userCoupon = &UserCoupon{userID, couponID, coupon.Name, coupon.Points, count, 0}
	}

	// update user coupon
	err = updateUserCoupon(stub, userCoupon)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("move coupon success")
	return shim.Success(nil)
}

// ===========================================================
// getUserCoupon
// ===========================================================
func (t *SimpleChaincode) getUserCoupon(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	queryString := "{\"selector\":{\"docType\":\"" + PREFIX_USER_COUPON + "\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===========================================================
// createOrder
// ===========================================================
func (t *SimpleChaincode) createOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// check args
	if len(args) != 4 {
		return shim.Error("Incorrect number of arguments. Expecting 4")
	}
	if len(args[0]) <= 0 {
		return shim.Error("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return shim.Error("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return shim.Error("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return shim.Error("4th argument must be a non-empty number")
	}

	id := args[0]
	userID := args[1]
	couponID := args[2]
	count, err := strconv.Atoi(args[3])
	if err != nil {
		return shim.Error("4th argument must be a number")
	}

	// check order existence
	var order *Order
	order, err = getOrderByID(stub, id)
	if err != nil {
		return shim.Error(err.Error())
	}
	if order != nil {
		return shim.Error("order already exists")
	}

	// check coupon
	userCoupon, err := getUserCouponByID(stub, userID, couponID)
	if err != nil {
		return shim.Error("Failed to get user coupon: " + err.Error())
	}
	if userCoupon == nil {
		return shim.Error("Failed to get user coupon")
	}
	if count > (userCoupon.Total - userCoupon.Used) {
		return shim.Error("Not enough coupon left")
	}

	// update user coupon
	userCoupon.Used = userCoupon.Used + count
	err = updateUserCoupon(stub, userCoupon)
	if err != nil {
		shim.Error(err.Error())
	}

	// create order
	order = &Order{PREFIX_ORDER, id, userID, userCoupon.Name, count, STATUS_LOCKED}
	err = createOrder(stub, order)
	if err != nil {
		shim.Error(err.Error())
	}

	fmt.Println("create order success")
	return shim.Success(nil)
}

// ===========================================================
// getOrder
// ===========================================================
func (t *SimpleChaincode) getOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	queryString := "{\"selector\":{\"docType\":\"" + PREFIX_ORDER + "\"}}"

	queryResults, err := getQueryResultForQueryString(stub, queryString)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

// ===========================================================
// auditOrder
// ===========================================================
func (t *SimpleChaincode) auditOrder(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	// check args
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	id := args[0]

	var order *Order
	order, err = getOrderByID(stub, id)
	if err != nil {
		return shim.Error(err.Error())
	}
	if order == nil {
		return shim.Error("cannot find this order")
	}

	order.Status = STATUS_AUDITED
	err = updateOrder(stub, order)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("audit order success")
	return shim.Success(nil)
}

// Deletes an entity from state
func (t *SimpleChaincode) deleteAll(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	key, err := stub.CreateCompositeKey(PREFIX_COUPON, nil)
	err = stub.DelState(key)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	key, err = stub.CreateCompositeKey(PREFIX_ORDER, nil)
	err = stub.DelState(key)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	key, err = stub.CreateCompositeKey(PREFIX_USER_COUPON, nil)
	err = stub.DelState(key)
	if err != nil {
		return shim.Error("Failed to delete state")
	}
	key, err = stub.CreateCompositeKey(PREFIX_USER_ORDER, nil)
	err = stub.DelState(key)
	if err != nil {
		return shim.Error("Failed to delete state")
	}

	return shim.Success(nil)
}

// ===========================================================
// below are private functions
// ===========================================================

// ===========================================================
// private getCouponByID
// ===========================================================
func getCouponByID(stub shim.ChaincodeStubInterface, id string) (*Coupon, error) {

	var err error

	key, err := stub.CreateCompositeKey(PREFIX_COUPON, []string{id})
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, nil
	}

	out := &Coupon{}
	err = json.Unmarshal([]byte(bytes), &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

// ===========================================================
// private getUserCouponByID
// ===========================================================
func getUserCouponByID(stub shim.ChaincodeStubInterface, userID string, couponID string) (*UserCoupon, error) {

	var err error

	key, err := stub.CreateCompositeKey(PREFIX_USER_COUPON, []string{userID, couponID})
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, nil
	}

	out := &UserCoupon{}
	err = json.Unmarshal([]byte(bytes), &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

// ===========================================================
// private getOrderByID
// ===========================================================
func getOrderByID(stub shim.ChaincodeStubInterface, id string) (*Order, error) {

	var err error

	key, err := stub.CreateCompositeKey(PREFIX_ORDER, []string{id})
	bytes, err := stub.GetState(key)
	if err != nil {
		return nil, err
	}

	if bytes == nil {
		return nil, nil
	}

	out := &Order{}
	err = json.Unmarshal([]byte(bytes), &out)
	if err != nil {
		return nil, err
	}

	return out, err
}

// ===========================================================
// private updateUserCoupon
// ===========================================================
func updateUserCoupon(stub shim.ChaincodeStubInterface, userCoupon *UserCoupon) error {

	var err error

	bytes, err := json.Marshal(userCoupon)
	if err != nil {
		return err
	}
	key, err := stub.CreateCompositeKey(PREFIX_USER_COUPON, []string{userCoupon.UserID, userCoupon.CouponID})
	err = stub.PutState(key, bytes)
	if err != nil {
		return err
	}

	return nil
}

// ===========================================================
// private updateCoupon
// ===========================================================
func updateCoupon(stub shim.ChaincodeStubInterface, coupon *Coupon) error {

	var err error

	bytes, err := json.Marshal(coupon)
	if err != nil {
		return err
	}
	key, err := stub.CreateCompositeKey(PREFIX_COUPON, []string{coupon.ID})
	err = stub.PutState(key, bytes)
	if err != nil {
		return err
	}

	return nil
}

// ===========================================================
// private updateOrder
// ===========================================================
func updateOrder(stub shim.ChaincodeStubInterface, order *Order) error {

	var err error

	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}
	key, err := stub.CreateCompositeKey(PREFIX_ORDER, []string{order.ID})
	err = stub.PutState(key, bytes)
	if err != nil {
		return err
	}

	return nil
}

// ===========================================================
// private createOrder
// ===========================================================
func createOrder(stub shim.ChaincodeStubInterface, order *Order) error {

	var err error

	bytes, err := json.Marshal(order)
	if err != nil {
		return err
	}
	key, err := stub.CreateCompositeKey(PREFIX_ORDER, []string{order.ID})
	err = stub.PutState(key, bytes)
	if err != nil {
		return err
	}

	return nil
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
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
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
