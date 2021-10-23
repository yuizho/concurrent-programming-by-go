package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sample1() {
	multiply := func(values []int, operand int) []int {
		vs := make([]int, len(values))
		for i, v := range values {
			vs[i] = v * operand
		}
		return vs
	}
	add := func(values []int, operand int) []int {
		vs := make([]int, len(values))
		for i, v := range values {
			vs[i] = v + operand
		}
		return vs
	}

	ints := []int{1, 2, 3, 4}
	// batch oriented pipeline
	for _, v := range add(multiply(ints, 2), 1) {
		fmt.Println(v)
	}
}

func sample2() {
	generator := func(done <-chan interface{}, integers ...int) <-chan int {
		intStream := make(chan int, len(integers))
		go func() {
			defer close(intStream)
			for _, i := range integers {
				select {
				case <-done:
					return
				case intStream <- i:
				}
			}
		}()
		return intStream
	}

	multiply := func(
		done <-chan interface{},
		intStream <-chan int,
		operand int,
	) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range intStream {
				select {
				case <-done:
					return
				case stream <- i * operand:
				}
			}
		}()
		return stream
	}
	add := func(
		done <-chan interface{},
		intStream <-chan int,
		operand int,
	) <-chan int {
		stream := make(chan int)
		go func() {
			defer close(stream)
			for i := range intStream {
				select {
				case <-done:
					return
				case stream <- i + operand:
				}
			}
		}()
		return stream
	}

	done := make(chan interface{})
	defer close(done)

	intStream := generator(done, 1, 2, 3, 4)
	pipeline := multiply(done, add(done, multiply(done, intStream, 2), 1), 2)

	for v := range pipeline {
		fmt.Println(v)
	}
}

func sample3() {
	repeat := func(
		done <-chan interface{},
		values ...interface{},
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for {
				for _, v := range values {
					select {
					case <-done:
						return
					case valueStream <- v:
					}
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	go func() {
		time.Sleep(1 * time.Microsecond)
		close(done)
	}()

	valueStream := repeat(done, 1, 2, 3, 4)
	for v := range valueStream {
		fmt.Print(v)
	}
}

func sample4() {
	repeatFn := func(
		done <-chan interface{},
		fn func() interface{},
		num int,
	) <-chan interface{} {
		valueStream := make(chan interface{})
		go func() {
			defer close(valueStream)
			for i := 0; i < num; i++ {
				select {
				case <-done:
					return
				case valueStream <- fn():
				}
			}
		}()
		return valueStream
	}

	done := make(chan interface{})
	defer close(done)

	rand := func() interface{} {
		return rand.Intn(1000000)
	}

	for num := range repeatFn(done, rand, 10) {
		fmt.Println(num)
	}
}

func main() {
	sample1()
	fmt.Println("-------------")
	sample2()
	fmt.Println("-------------")
	sample3()
	fmt.Println("\n-------------")
	sample4()
}
