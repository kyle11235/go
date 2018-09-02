package main

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type SimpleChaincode struct {
}

// coupon
type Coupon struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Points int    `json:"points"`
	Total  int    `json:"total"`
	Used   int    `json:"used"`
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

// prefix
const (
	PREFIX_COUPON      = "coupon"
	PREFIX_USER_COUPON = "user-coupon"
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
	}

	// result
	message := "invoke did not find func: " + function
	fmt.Println(message)
	return shim.Error(message)
}

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
		return shim.Error("4rd argument must be a number")
	}

	// check existence
	compositeKey, err := stub.CreateCompositeKey(PREFIX_COUPON, []string{id})
	couponBytes, err := stub.GetState(compositeKey)
	if err != nil {
		return shim.Error("Failed to get coupon: " + err.Error())
	} else if couponBytes != nil {
		fmt.Println("This coupon already exists: " + id)
		return shim.Error("This coupon already exists: " + id)
	}

	// create
	coupon := &Coupon{id, name, points, count, 0}
	newCouponBytes, err := json.Marshal(coupon)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(compositeKey, newCouponBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	fmt.Println("create coupon success")
	return shim.Success(nil)
}

// ===========================================================
func (t *SimpleChaincode) getCoupon(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error
	iterator, err := stub.GetStateByPartialCompositeKey(PREFIX_COUPON, nil)
	if err != nil {
		return shim.Error(err.Error())
	}
	if iterator == nil {
		return shim.Error("getCoupon error")
	}
	defer iterator.Close()

	list := CouponList{make([]Coupon, 0)}
	for iterator.HasNext() {
		next, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		item := Coupon{}
		itemBytes := next.GetValue()
		if itemBytes != nil {
			err = json.Unmarshal(itemBytes, &item)
		}
		list.Items = append(list.Items, item)
	}
	listBytes, err := json.Marshal(list)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("get coupon success")
	return shim.Success(listBytes)
}

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
	coupon := Coupon{}

	couponKey, err := stub.CreateCompositeKey(PREFIX_COUPON, []string{couponID})
	couponBytes, err := stub.GetState(couponKey)
	if err != nil {
		return shim.Error("Failed to get coupon: " + err.Error())
	}
	if couponBytes == nil {
		fmt.Println("This coupon NOT exists: " + couponID)
		return shim.Error("This coupon NOT exists: " + couponID)
	}
	err = json.Unmarshal([]byte(couponBytes), &coupon)
	if err != nil {
		return shim.Error(err.Error())
	}

	// update coupon
	coupon.Used = coupon.Used + count
	newCouponBytes, err := json.Marshal(coupon)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(couponKey, newCouponBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	// get user coupon
	userCoupon := UserCoupon{}

	userCouponKey, err := stub.CreateCompositeKey(PREFIX_USER_COUPON, []string{userID, couponID})
	userCouponBytes, err := stub.GetState(userCouponKey)
	if err != nil {
		return shim.Error("Failed to get user coupon: " + err.Error())
	}
	if userCouponBytes != nil {
		// update
		err = json.Unmarshal([]byte(userCouponBytes), &userCoupon)
		if err != nil {
			return shim.Error(err.Error())
		}
		userCoupon.Total = userCoupon.Total + count
	} else {
		// create
		userCoupon.UserID = userID
		userCoupon.CouponID = couponID
		userCoupon.Points = coupon.Points
		userCoupon.Total = count
		userCoupon.Used = 0
	}

	// update user coupon
	newUserCouponBytes, err := json.Marshal(userCoupon)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(userCouponKey, newUserCouponBytes)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("move coupon success")
	return shim.Success(nil)
}

// ===========================================================
func (t *SimpleChaincode) getUserCoupon(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	userID := args[0]

	iterator, err := stub.GetStateByPartialCompositeKey(PREFIX_USER_COUPON, []string{userID})
	if err != nil {
		return shim.Error(err.Error())
	}
	if iterator == nil {
		return shim.Error("getUserCoupon error")
	}
	defer iterator.Close()

	list := UserCouponList{make([]UserCoupon, 0)}
	for iterator.HasNext() {
		next, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		item := UserCoupon{}
		itemBytes := next.GetValue()
		if itemBytes != nil {
			err = json.Unmarshal(itemBytes, &item)
		}
		list.Items = append(list.Items, item)
	}
	listBytes, err := json.Marshal(list)
	if err != nil {
		return shim.Error(err.Error())
	}

	fmt.Println("get user coupon success")
	return shim.Success(listBytes)
}
