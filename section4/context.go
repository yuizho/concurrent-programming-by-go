package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func cancelSample() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		for {
			fmt.Println("Waiting for time out by context.WithCancel")

			select {
			case <-ctx.Done():
				fmt.Println("Exit now!")
				wg.Done()
				return
			default:
			}
		}
	}()

	go func() {
		time.Sleep(1 * time.Microsecond)
		cancel()
	}()

	wg.Wait()
}

func timeoutSample() {
	var wg sync.WaitGroup
	wg.Add(1)

	ctx := context.Background()

	go func(ctx context.Context) {
		// wrap passed context
		innerCtx, cancel := context.WithTimeout(ctx, 500*time.Nanosecond)
		defer cancel()
		for {
			fmt.Println("Waiting for time out by context.WithTimeout")

			select {
			case <-innerCtx.Done():
				fmt.Println("Exit now!")
				wg.Done()
				return
			default:
			}
		}
	}(ctx)

	wg.Wait()
}

func keyValSample() {
	// define custome type to prevent key confliction
	// when a key type of map is different,  go lang judge that these keys are different one.
	// even if the value is same.
	type ctxKey int

	const (
		ctxUserId ctxKey = iota
		ctxAuthToken
	)

	UserId := func(ctx context.Context) string {
		return ctx.Value(ctxUserId).(string)
	}

	AuthToken := func(ctx context.Context) string {
		return ctx.Value(ctxAuthToken).(string)
	}

	handleResponse := func(ctx context.Context) {
		fmt.Printf(
			"handling response for %v (%v)\n",
			// hide the keys from context client
			UserId(ctx),
			AuthToken(ctx),
		)
	}

	processRequest := func(
		userId, authToken string,
	) {
		ctx := context.WithValue(context.Background(), ctxUserId, userId)
		ctx = context.WithValue(ctx, ctxAuthToken, authToken)
		handleResponse(ctx)
	}

	processRequest("yui", "abc123")
}

func main() {
	cancelSample()
	fmt.Println("--------------")
	timeoutSample()
	fmt.Println("--------------")
	keyValSample()
}
