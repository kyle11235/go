package main

import "fmt"

// channel - https://www.jtolio.com/2016/03/go-channels-are-bad-and-bou-should-feel-bad/
func main() {
	// ch <- v    // Send v to channel ch.
	// v := <-ch  // Receive from ch, and assign value to v.

	fib := func(n int, ch chan int) {
		a, b := 0, 1
		for i := 0; i < n; i++ {
			ch <- a
			a, b = b, a+b
		}
		close(c)
	}

	ch1 := make(chan int)
	go fib(6, ch1)
	for i := range ch1 {
		fmt.Println(i) // receive blocks until the send side is ready
	}
}
