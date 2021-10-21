package main

import (
	"fmt"
	"time"
)

func sample1() {
	start := time.Now()
	c := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Second)
		close(c)
	}()

	fmt.Println("Blocking on read...")
	select {
	case <-c:
		fmt.Printf("Unblocked %v later.\n", time.Since(start))
	}
}

func sample2() {
	c1 := make(chan interface{})
	close(c1)
	c2 := make(chan interface{})
	close(c2)

	var c1Count, c2Count int
	for i := 1000; i >= 0; i-- {
		select {
		case <-c1:
			c1Count++
		case <-c2:
			c2Count++
		}
	}

	fmt.Printf("c1Count: %d\nc2Count: %d\n", c1Count, c2Count)
}

func sample3() {
	var c <-chan int
	select {
	case <-c:
		// this channel is blocked forever
	case <-time.After(1 * time.Second):
		// this block will be called in a second
		fmt.Println("Timed out.")
	}
}

func sample4() {
	start := time.Now()
	var c1, c2 <-chan int
	select {
	case <-c1:
		// this channel is blocked forever
	case <-c2:
		// this channel is blocked forever
	default:
		// default block will be called , when all channels are blocked
		// that's why "default" is used with loop frequently
		fmt.Printf("In defalt after %v\n\n", time.Since(start))
	}
}

func sample5() {
	done := make(chan interface{})
	go func() {
		time.Sleep(3 * time.Second)
		close(done)
	}()

	workCounter := 0
loop:
	for {
		select {
		case <-done:
			break loop
		default:
			fmt.Print("*")
		}

		workCounter++
		time.Sleep(500 * time.Millisecond)
	}

	fmt.Printf("\nAchieved %v cycles of work before signalled to stop.\n", workCounter)
}

func main() {
	sample1()
	fmt.Println("-------------------")
	sample2()
	fmt.Println("-------------------")
	sample3()
	fmt.Println("-------------------")
	sample4()
	fmt.Println("-------------------")
	sample5()
}
