package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func ByAtomc() {
	var count int64
	var wg sync.WaitGroup
	wg.Add(1000000)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		go func() {
			atomic.AddInt64(&count, 1)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("atomc: %v spent\n", time.Since(start))
	fmt.Printf("the count is %v\n", count)
}

func ByMutex() {
	var count int64
	var lock sync.Mutex
	var wg sync.WaitGroup
	wg.Add(1000000)
	start := time.Now()
	for i := 0; i < 1000000; i++ {
		go func() {
			lock.Lock()
			count += 1
			defer lock.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Printf("mutex: %v spent\n", time.Since(start))
	fmt.Printf("the count is %v\n", count)
}

func main() {
	ByAtomc()
	fmt.Println("--------------")
	ByMutex()
}
