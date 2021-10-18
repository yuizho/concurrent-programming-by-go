package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := sync.NewCond(&sync.Mutex{})
	// a variable shared to every groutine
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		// this lock is for pop queue
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue")
		c.L.Unlock()
		// send signal to the goroutine which is waiting for the longest time.
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		fmt.Printf("this loop id is %v\n", i)
		// this lock is for adding queue
		c.L.Lock()
		for len(queue) == 2 {
			// block main goroutine until sent signal
			// when the queu length is not 2, this loop continues checking queue length
			fmt.Printf("%v is blocked\n", i)
			c.Wait()
		}
		fmt.Println("Adding to queue")
		queue = append(queue, struct{}{})
		// pop queue in 1 second
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
	fmt.Printf("the queue length is %v\n", len(queue))
}
