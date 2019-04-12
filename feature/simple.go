package main

import (
	"fmt"
)

var i, java, nodejs, golang = 10, false, false, true

func swap(x, y string) (string, string) {
	return y, x
}

func main() {

	j := 20 
	fmt.Println(i, j, java, nodejs, golang)

	a, b := swap("a", "b")
	fmt.Println(a, b)

}
