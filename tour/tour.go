package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strings"
	"time"
)

// variable
var i, java, node, golang = 100, false, false, true

// Vertex is struct
type Vertex struct {
	X float64
	Y float64
}

// function
func add(x int, y int) int {
	return x + y
}

func swap(x, y string) (string, string) {
	return y, x
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}

// method - function on type, with value as receiver
func (v Vertex) test1() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// Tester is interface - set of methods
type Tester interface {
	test2() float64
}

// implementation - method of same signature, with pointer as reveiver
func (v *Vertex) test2() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// MyError is my error
type MyError struct {
	When time.Time
	What string
}

// MyError implements Error method of error interface
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

// return error interface
// same with errors.New("my error")
func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {

	// 1. func
	fmt.Println(add(1, 2)) // 3

	// return
	a, b := swap("a", "b")
	fmt.Println(a, b) // b a

	// 2. variable
	var j = "aString"
	fmt.Println(i, j, java, node, golang) // 100 aString false false true

	// implicit variable
	d := "d"
	fmt.Println(d) // d

	// type conversion
	f := float64(i)
	fmt.Println(f) // 100

	// 3. for loop, i only valid inside loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(i)   // 100
	fmt.Println(sum) // 45

	// while loop
	for sum < 200 {
		sum += sum
	}
	fmt.Println(sum) // 360

	// switch
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.", os) // OS X.
	}

	// 4. defer
	defer fmt.Println("world") // this will run after main is returned
	fmt.Println("hello")       // hello world

	// 5. pointer
	var p1 *int
	p1 = &i

	fmt.Println(*p1) // 100
	*p1 = 21
	fmt.Println(i) // 21

	// 6. struct
	var (
		v1 = Vertex{1, 2}  // has type Vertex
		v2 = Vertex{X: 1}  // Y:0 is implicit
		v3 = Vertex{}      // X:0 and Y:0
		p2 = &Vertex{1, 2} // has type *Vertex
	)
	p2.X = 4
	fmt.Println(v1, v2, v3, *p2, p2.X)
	// {1 2} {1 0} {0 0} {4 2} 4

	// 7. array - static length

	// 1) empty array
	var arr1 [2]string
	fmt.Println(arr1) // []

	// 2) literal array
	arr2 := [2]string{"aaa", "bbb"}
	fmt.Println(arr2) // [aaa bbb]

	// 3) use index
	arr1[0] = "Hello"
	arr1[1] = "World"
	fmt.Println(arr1[0], arr1[1]) // Hello World
	fmt.Println(arr1)             // [Hello World]

	// 4) range over
	for i, v := range arr1 {
		fmt.Printf("index%v = %v\n", i, v)
	}

	// 5) for loop
	for i := 0; i < len(arr1); i++ {
		fmt.Printf("index%d = %v\n", i, arr1[i])
	}

	// 8. slice -  dynamic length, a view/reference to an array. like java arraylist

	// 1) nil slice
	var s1 []int
	printSlice(s1) // len=0 cap=0 []
	if s1 == nil {
		fmt.Println("nil!") // nil!
	}

	// 2) literal slice
	s2 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s2) // [0 1 2 3 4 5]

	// 3) use append, new array is allocated
	s1 = append(s1, 1, 1)
	printSlice(s1) // len=2 cap=2 [1 1]

	// 4) use make, know how much you need
	s3 := make([]int, 5)    // len = 5, cap = 5
	s4 := make([]int, 0, 5) // len = 0, cap = 5
	fmt.Println(s3, s4)     // [0 0 0 0 0] []

	// 5) half open range
	arr3 := [6]int{0, 1, 2, 3, 4, 5}
	s5 := arr3[1:4] // half open range
	fmt.Println(s5) // [1 2 3]

	// 6) length of slice vs capacity of array
	s6 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s6) // len=6 cap=6 [2 3 5 7 11 13]

	s6 = s6[:0]    // Slice the slice to give it zero length.
	printSlice(s6) // len=0 cap=6 []

	s6 = s6[:4]    // Extend its length.
	printSlice(s6) // len=4 cap=6 [2 3 5 7]

	s6 = s6[2:]    // Drop its first two values.
	printSlice(s6) // len=2 cap=4 [5 7]

	// 7) range over
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("index%d = %d\n", i, v)
	}

	// 8) for loop
	for i := 0; i < len(pow); i++ {
		fmt.Printf("index%d = %v\n", i, pow[i])
	}

	// 9. map

	// 1) nil map
	var m1 map[string]int
	fmt.Println(m1) // map[]
	if m1 == nil {
		fmt.Println("nil!") // nil!
	}

	// 2) literal map
	var m2 = map[string]Vertex{
		"Bell Labs": {10, 20},
		"Google":    {30, 40},
	}
	fmt.Println(m2) // map[Bell Labs:{10 20} Google:{30 40}]

	// 3) use make
	m1 = make(map[string]int)
	m1["key"] = 42
	fmt.Printf("value=%v\n", m1["key"]) // value=42

	delete(m1, "key")
	fmt.Printf("value=%v\n", m1["key"]) // value=0

	v, exist := m1["key"]
	fmt.Printf("value=%v, exist=%v\n", v, exist) // value=0, exist=false

	// 4) range over
	for k, v := range m2 {
		fmt.Printf("%v, %v\n", k, v)
	}
	// Bell Labs, {10 20}
	// Google, {30 40}

	// 10. method - function on type
	v4 := Vertex{3, 4}
	fmt.Println(v4.test1()) // 5

	// 11. interface - data/method/interface/implementation are decoupled with object/duck typing(structual typing actually)
	var tester Tester = &v4
	fmt.Printf("%v, %T\n", tester.test2(), tester) // 5, *main.Vertex - the type of the interface variable

	// 12. error interface
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// 13. reader
	r := strings.NewReader("Hello, Reader!")
	bytes := make([]byte, 8)
	for {
		n, err := r.Read(bytes)
		fmt.Printf("n = %v err = %v bytes = %v\n", n, err, bytes)
		fmt.Printf("bytes[:n] = %q\n", bytes[:n])
		if err == io.EOF {
			break
		}
	}
	/*
		n = 8 err = <nil> bytes = [72 101 108 108 111 44 32 82]
		bytes[:n] = "Hello, R"

		n = 6 err = <nil> bytes = [101 97 100 101 114 33 32 82]
		bytes[:n] = "eader!"

		n = 0 err = EOF bytes = [101 97 100 101 114 33 32 82]
		bytes[:n] = ""
	*/

	// 14. goroutine - lightweight thread (not actually is) managed by the Go runtime
	say := func(s string) {
		for i := 0; i < 2; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(s)
		}
	}
	go say("in routine") // start routine
	say("in main")       // keep main

	// 15. encoding
	/*
		1. go source code is utf8
		2. There are two places in the language that Go does do UTF-8 decoding of strings for you.
			when you do for i, r := range s the r is a Unicode code point as a value of type rune
			when you do the conversion []rune(s), Go decodes the whole string to runes
	*/

}
