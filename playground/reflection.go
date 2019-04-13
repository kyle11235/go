package main

import (
	"fmt"
	"reflect"
)

type Reader interface {
	Read(p []byte) (n int, err error)
}

// r's interface type is Reader
var r Reader

// i's interface type is interface{} - empty set of methods
var i interface{}

// https://blog.golang.org/laws-of-reflection
func main() {

	// 1. At the basic level, reflection is just a mechanism to examine the type and value pair stored inside an interface variable
	var x float64 = 3.4
	fmt.Println("type:", reflect.TypeOf(x)) // type: float64

	v := reflect.ValueOf(x)
	fmt.Println("type:", v.Type())                               // float64
	fmt.Println("kind is float64:", v.Kind() == reflect.Float64) // true
	fmt.Println("value:", v.Float())                             //3.4

	// 2. Given a reflect.Value we can recover an interface value using the Interface method;

}

