package main

import "fmt"

// function as value
func compute(value float64, fn func(float64) float64) float64 {
	return fn(value)
}

// function type
type Func func(float64) float64

func compute1(value float64, fn Func) float64 {
	return fn(value)
}

func main() {

	decrease := func(x float64) float64 {
		return x - 1
	}
	fmt.Println(compute(5, decrease)) // 4

	var increase Func = func(x float64) float64 {
		return x + 1
	}
	fmt.Println(compute1(5, increase)) // 6

}
