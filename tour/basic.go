package main

import (
	"fmt"
	"io"
	"math"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"time"
)

// variable
var java, golang, nodejs bool

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

// function as value
func compute(value float64, fn func(float64) float64) float64 {
	return fn(value)
}

// closure - funcion as return value
func adder() func(int) int {
	sum := 0 // this is variable in the closure
	return func(x int) int {
		sum += x
		return sum
	}
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
	var i, j = 100, "golang"
	fmt.Println(c, python, java, i, j) // false false false 100 golang

	// implicit variable
	d := "d"
	fmt.Println(d) // d

	// type conversion
	f := float64(i)
	fmt.Println(f) // 100

	// 3. for loop
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum) // 180

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

	// 7. array
	var arr1 [2]string
	arr1[0] = "Hello"
	arr1[1] = "World"
	fmt.Println(arr1[0], arr1[1]) // Hello World
	fmt.Println(arr1)             // [Hello World]

	// literal array
	arr1 = [2]string{"Hello", "World"}
	fmt.Println(arr1) // [Hello World]

	// 8. slice -  a view/reference to an array. like java arraylist

	// 1) make
	s3 := make([]int, 5)    // len = 5, cap = 5
	s4 := make([]int, 0, 5) // len = 0, cap = 5
	fmt.Println(s3, s4)     // [0 0 0 0 0] []

	// 2) half open range
	arr2 := [6]int{0, 1, 2, 3, 4, 5}
	s2 := arr2[1:4] // half open range
	fmt.Println(s2) // [1 2 3]

	// 3) literal slice
	s1 := []int{0, 1, 2, 3, 4, 5}
	fmt.Println(s1) // [0 1 2 3 4 5]

	// 4) length of slice vs capacity of array
	s5 := []int{2, 3, 5, 7, 11, 13}
	printSlice(s5) // len=6 cap=6 [2 3 5 7 11 13]

	s5 = s5[:0]    // Slice the slice to give it zero length.
	printSlice(s5) // len=0 cap=6 []

	s5 = s5[:4]    // Extend its length.
	printSlice(s5) // len=4 cap=6 [2 3 5 7]

	s5 = s5[2:]    // Drop its first two values.
	printSlice(s5) // len=2 cap=4 [5 7]

	// 5) nil slice
	var s6 []int
	fmt.Println(s6, len(s6), cap(s6)) // [] 0 0
	if s6 == nil {
		fmt.Println("nil!") // nil!
	}

	// 6) append to slice, new array is allocated
	s6 = append(s6, 2, 3)
	printSlice(s6) // len=3 cap=4 [6 6 6]

	// 9. range over slice
	var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}
	for i, v := range pow {
		fmt.Printf("index%d = %d\n", i, v)
	}

	// 10. map
	m1 := make(map[string]Vertex)
	m1["Bell Labs"] = Vertex{
		10, 20, // comma is needed
	}
	fmt.Println(m1["Bell Labs"]) // {10 20}

	// literal map
	var m2 = map[string]Vertex{
		"Bell Labs": {10, 20},
		"Google":    {30, 40},
	}
	fmt.Println(m2) // map[Bell Labs:{10 20} Google:{30 40}]

	// operate map
	m3 := make(map[string]int)
	m3["key"] = 42
	fmt.Println("The value:", m3["key"]) // 42

	delete(m3, "key")
	fmt.Println("The value:", m3["key"]) // 0

	v, exist := m3["key"]
	fmt.Println("The value:", v, "Exist?", exist) // false

	// range over map
	for k, v := range m2 {
		fmt.Printf("%v, %v\n", k, v)
	}
	// Bell Labs, {10 20}
	// Google, {30 40}

	// 11. function as value
	minus := func(x float64) float64 {
		return x - 1
	}
	fmt.Println(compute(5, minus)) // 4

	// 12. function closure
	myadder := adder()
	fmt.Println(
		myadder(1),
		myadder(2),
	)
	// 1 3

	// 13. method - function on type
	v4 := Vertex{3, 4}
	fmt.Println(v4.test1()) // 5

	// 14. interface - declaration and implementation of interface are decoupled with object/duck typing
	var tester Tester = &v4
	fmt.Printf("%v, %T\n", tester.test2(), tester) // 5, *main.Vertex - the type of the interface variable

	// 15. error interface
	if err := run(); err != nil {
		fmt.Println(err)
	}

	// 16. reader
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

	// 17. goroutine - lightweight thread (not actually is) managed by the Go runtime
	say := func(s string) {
		for i := 0; i < 2; i++ {
			time.Sleep(1 * time.Second)
			fmt.Println(s)
		}
	}
	go say("in routine") // start routine
	say("in main")       // keep main

	// 18. channel - https://www.jtolio.com/2016/03/go-channels-are-bad-and-you-should-feel-bad/
	// ch <-v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and assign value to v.

	fibonacci := func(n int, c chan int) {
		x, y := 0, 1
		for i := 0; i < n; i++ {
			c <- x
			x, y = y, x+y
		}
		close(c)
	}

	c1 := make(chan int)
	go fibonacci(6, c1)
	for i := range c1 {
		fmt.Println(i) // receive blocks until the send side is ready
	}

	// 19. sync.Mutex
	var count = 0
	var lock sync.Mutex
	produce := func() {
		for {
			lock.Lock()
			if count == 0 {
				count++
				fmt.Println("produce, count=" + strconv.Itoa(count))
			}
			lock.Unlock()
			time.Sleep(1 * time.Second)
		}
	}
	consume := func() {
		for {
			lock.Lock()
			if count == 1 {
				count--
				fmt.Println("consume, count=" + strconv.Itoa(count))
			}
			lock.Unlock()
			time.Sleep(1 * time.Second)
		}
	}

	go produce()
	go produce()
	go consume()
	go consume()
	// produce, count=1
	// consume, count=0
	// ...

	time.Sleep(1000 * time.Second) // keep main

}
