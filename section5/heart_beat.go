package main

import (
	"fmt"
	"time"
)

func doWork(
	done <-chan interface{},
	pluseInterval time.Duration,
) (<-chan interface{}, <-chan time.Time) {
	heartbeat := make(chan interface{})
	results := make(chan time.Time)
	go func() {
		defer close(heartbeat)
		defer close(results)

		// a ticker for heart beat
		// tick return chan Time
		pluse := time.Tick(pluseInterval)
		// a ticker for some task
		workGen := time.Tick(2 * pluseInterval)

		sendPulse := func() {
			select {
			case heartbeat <- struct{}{}:
			default:
			}
		}
		// send task result and heart beat
		sendResult := func(r time.Time) {
			for {
				select {
				case <-done:
					return
				case <-pluse:
					sendPulse()
				case results <- r:
					return
				}
			}
		}

		for {
			select {
			case <-done:
				return
			case <-pluse:
				sendPulse()
			case r := <-workGen:
				sendResult(r)
			}
		}
	}()
	return heartbeat, results
}

func main() {
	done := make(chan interface{})
	time.AfterFunc(10*time.Second, func() { close(done) })

	const timeout = 2 * time.Second
	// divide timeout sec to reduce heatbeat frequency
	heartbeat, results := doWork(done, timeout/2)
	for {
		select {
		case _, ok := <-heartbeat:
			if !ok {
				return
			}
			fmt.Println("pluse")
		case r, ok := <-results:
			if ok == false {
				return
			}
			fmt.Printf("results %v\n", r.Second())

		case <-time.After(timeout):
			// when there is no other signal, this block will be executed
			// if this block is implemented by default, the for loop is finished immediately
			fmt.Println("timeout")
			return
		}
	}
}
