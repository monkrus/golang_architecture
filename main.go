package main

import (
	"context"
	"time"
)

func main() {
//PARENT context
	ctx := context.Background()
// CHILD Context with timeout returns two values
//waits for half second
	ctx, cancelF := context.WithTimeout(ctx, 500*time.Millisecond)
// cleanup, shutting it down
	defer cancelF()
// then some work starts
	time.Sleep(100* time.Millisecond)
//check the status of the channel

//when done channel gets closed it is in `"false" and has a ZERO value.`
//what is coming off the channel?
select {

case <-ctx.Done();
//not finished because done returned a VALUE !
	fmt.Printl("work did not finished")
default:
	fmt.Printl("work  finished")
	}

}

