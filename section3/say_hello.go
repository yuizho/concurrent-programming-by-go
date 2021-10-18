package main

import (
	"fmt"
	"sync"
)

func v1SayHello() {
	var wg sync.WaitGroup
	sayHello := func() {
		defer wg.Done()
		fmt.Println("hello!")
	}

	wg.Add(1)
	go sayHello()
	wg.Wait()
}

func v2SayHello() {
	var wg sync.WaitGroup
	salutation := "hello"
	wg.Add(1)
	go func() {
		defer wg.Done()
		// closure can access to enclosed variable
		// even if the function is groutine!!!!
		salutation = "welcome"
	}()
	wg.Wait()
	fmt.Println(salutation)
}

func v3SayHello() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "gretings", "good day"} {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// the all of greeting will be good day.....
			// the loop is finished before creating goroutines.
			// groutine moves the salutation variable to heap,
			// that's why every groutine  could reference the salutation variable and all of greeting would be good day.
			fmt.Println(salutation)
		}()
	}
	wg.Wait()
}

func v4SayHello() {
	var wg sync.WaitGroup
	for _, salutation := range []string{"hello", "gretings", "good day"} {
		wg.Add(1)
		// this version resolves v3 problem by parameter (value copy)
		go func(salutation string) {
			defer wg.Done()
			fmt.Println(salutation)
		}(salutation)
	}
	wg.Wait()
}

func main() {
	v1SayHello()
	v2SayHello()
	v3SayHello()
	v4SayHello()
}
