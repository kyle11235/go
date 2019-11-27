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

// Reflection is the ability to check types; it's a form of metaprogramming
// interface{} = unknown type/every type's interface
func main() {

	// 1. use as unknown type
	// e.g. fmt.Printf(format string, a ...interface{}) takes any number of arguments of type interface{}.

	var i interface{}
	describe(i) // (<nil>, <nil>)

	i = "hello"
	describe(i) // (hello, string)

	i = &T{"hello"}
	describe(i) // (&{hello}, *main.T)

	// 2. check type with reflect
	i = 3.4
	reflectType := reflect.TypeOf(i)
	fmt.Printf("type=%v\n", reflectType) // type=float64

	reflectValue := reflect.ValueOf(i)
	fmt.Printf("type=%v\n", reflectValue.Type())                               // type=float64
	fmt.Printf("kind is float64=%v\n", reflectValue.Kind() == reflect.Float64) // kind is float64=true
	fmt.Printf("value=%v\n", reflectValue.Float())                             // value=3.4

	// 3. recover type, check fmt.Println -> p.doPrintln -> p.printArg -> switch f := arg.(type)
	// to pointer
	i = &T{"hello"}
	if t, ok := i.(*T); ok {
		fmt.Printf("type=%T\n", t)    // type=*main.T
		fmt.Printf("value=%v\n", t.S) // value=hello
	}

	// to int
	i = 100
	if t, ok := i.(int); ok {
		fmt.Printf("type=%T\n", t)  // type=int
		fmt.Printf("value=%v\n", t) // value=100
	}

	// to func
	var fn Func = func(s string) {
		fmt.Println(s)
	}

	i = fn
	if fn, ok := i.(Func); ok {
		fn("world") // world
	}

}
