package main

import (
	"fmt"
	"sync"
)

// A memory model allows a compiler to perform many important optimizations
// which influence the order of read and write operations of potentially shared variables
// so sync/serialize/lock to prevent race condition/access to shared data

func main() {

	// 1. channel
	// ...

	// 2. lock
	var l sync.Mutex
	var s string

	f := func() {
		s = "hello, world"
		l.Unlock()
	}

	l.Lock()
	go f()
	l.Lock()
	l.Unlock()
	fmt.Println(s) // hello, world is set for sure

	// 3. race condition
	var a, b int

	x := func() {
		a = 1 // set a first
		b = 1
	}

	y := func() {
		fmt.Println(b) // print b first
		fmt.Println(a)
	}
	go x()
	y() // it can happen 1 0, because order among multiple goroutines are not guaranteed

}

// enable race condition dector, detect unsynchronized accesses to shared variables 
// go run -race sync.go

// hello, world
// 0
// ==================
// WARNING: DATA RACE
// Read at 0x00c00008e020 by main goroutine:
//   main.main.func3()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:40 +0xba
//   main.main()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:43 +0x359

// Previous write at 0x00c00008e020 by goroutine 7:
//   main.main.func2()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:34 +0x4f

// Goroutine 7 (running) created at:
//   main.main()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:42 +0x347
// ==================
// 1
// ==================
// WARNING: DATA RACE
// Write at 0x00c00008e028 by goroutine 7:
//   main.main.func2()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:35 +0x69

// Previous read at 0x00c00008e028 by main goroutine:
//   main.main.func3()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:39 +0x53
//   main.main()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:43 +0x359

// Goroutine 7 (running) created at:
//   main.main()
//       /Users/kyle/go/src/github.com/kyle11235/go/memory_model/sync.go:42 +0x347
// ==================
// Found 2 data race(s)
// exit status 66