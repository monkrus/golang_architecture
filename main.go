package main

import (
	"context"
	"fmt"
	"runtime"
	"time"
)

func main() {
	fmt.Printf("GOROUTINES RUNNING %d\n", runtime.NumGoroutine())
	//PARENT context
	ctx := context.Background()
	// CHILD Context with timeout returns two values
	ctx, cancelF := context.WithCancel(ctx)
	// cleanup, shutting it down
	defer cancelF()

	// start a bunch of goroutines
	for i := 0; i < 100; i++ {
		// call an anonymous func
		go func(n int) {
			fmt.Println("launching goroutine", n)
			for {
				select {
				case <-ctx.Done():
					runtime.Goexit()
				// `return` is an alternative to `Goexit`
				//return
				default:
					fmt.Printf("goroutine %d is doing work\n", n)
					time.Sleep(50 * time.Millisecond)
				}
			}
		}(i)

	}
	time.Sleep(time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER ONE MILLISECOND %d\n", runtime.NumGoroutine())
	cancelF()
	time.Sleep(100 * time.Millisecond)
	fmt.Printf("GOROUTINES RUNNING AFTER Cancelfunc called %d\n", runtime.NumGoroutine())

}
