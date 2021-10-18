package main

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})

	var token struct{}
	sender := func() {
		defer wg.Done()
		// wait until start
		<-begin
		for i := 0; i < b.N; i++ {
			// send message to the receiver goroutine
			c <- token
		}
	}
	receiver := func() {
		defer wg.Done()
		// wait until start
		<-begin
		for i := 0; i < b.N; i++ {
			// receive the message sent by the sender groutine (and do nothing)
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()
	b.StartTimer()
	close(begin) // start
	wg.Wait()
}
