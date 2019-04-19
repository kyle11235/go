package main

import (
	"fmt"
	"reflect"
)

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M() {
	fmt.Println(t.S)
}

func describe(i interface{}) {
	fmt.Printf("(%v, %T)\n", i, i)
}

type Func func(s string)

// Reflection is the ability to examine types; it's a form of metaprogramming
func main() {

	// 1. empty interface are used by code that handles values of unknown type. For example, fmt.Print takes any number of arguments of type interface{}.
	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = "hello"
	describe(i) // (hello, string)

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)

	// 2. examine the type and value, float64's interface is interface{}, which is accepted by reflect.TypeOf(i interface{})
	var x float64 = 3.4
	reflectType := reflect.TypeOf(x)
	fmt.Println("type:", reflectType) // type: float64

	reflectValue := reflect.ValueOf(x)
	fmt.Println("type:", reflectValue.Type())                               // float64
	fmt.Println("kind is float64:", reflectValue.Kind() == reflect.Float64) // true
	fmt.Println("value:", reflectValue.Float())                             // 3.4

	// 3. recover the interface value
	var fn Func = func(s string){
		fmt.Println(s) 
	}

	i = fn
	if fn, ok := i.(Func); ok {
		fn("world") // world
	}

}
