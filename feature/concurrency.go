

package main

import (
	"fmt"
	"sync"
	"strconv"
	"time"
)

// goroutine - multiplex independently executing functions—coroutines—onto a set of threads
func main(){

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

