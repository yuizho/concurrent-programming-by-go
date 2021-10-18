package main

import (
	"fmt"
	"sync"
)

func main() {
	var count int
	var lock sync.Mutex

	increment := func() {
		lock.Lock()
		defer lock.Unlock()
		count++
		fmt.Printf("Incrementing: %d\n", count)
	}

	decrement := func() {
		lock.Lock()
		defer lock.Unlock()
		count--
		fmt.Printf("Decrementing: %d\n", count)
	}

	var arthmetic sync.WaitGroup

	for i := 0; i <= 5; i++ {
		arthmetic.Add(1)
		go func() {
			defer arthmetic.Done()
			increment()
		}()
	}
	for i := 0; i <= 5; i++ {
		arthmetic.Add(1)
		go func() {
			defer arthmetic.Done()
			decrement()
		}()
	}

	arthmetic.Wait()

	fmt.Println("Arthmetic complete.")
}
