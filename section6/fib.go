package main

import (
	"fmt"
	"time"
)

func simpleFib() {
	var fib func(n int) int
	fib = func(n int) int {
		if n <= 2 {
			return 1
		}
		return fib(n-1) + fib(n-2)
	}

	now := time.Now()
	fmt.Printf("simple fib(10) = %d\n", fib(10))
	fmt.Printf("%v\n", time.Since(now))
}

func fibWithDeque() {
	var fib func(n int) <-chan int
	fib = func(n int) <-chan int {
		result := make(chan int)
		go func() {
			defer close(result)
			if n <= 2 {
				result <- 1
				return
			}
			result <- <-fib(n-1) + <-fib(n-2)
		}()
		return result
	}

	now := time.Now()
	fmt.Printf("deque fib(10) = %d\n", <-fib(10))
	fmt.Printf("%v\n", time.Since(now))
}

func main() {
	fibWithDeque()
	fmt.Println("----------")
	simpleFib()
}
