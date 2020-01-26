package main

import (
"fmt"
"context"
)

func main() {
// parent context
ctx:= context.WithValue(context.Background(), "userID", 12345)
// child context
ctx = context.WithValue(ctx, 1, "admin") 
// printout
if v :=ctx.Value("userID"); v !=nil {
	fmt.Println(v)
} else {
fmt.Println("no value associated with that key")
}

if v :=ctx.Value(1); v !=nil {
	fmt.Println(v)
} else {
	fmt.Println("no value associated with that key")
}

if v := ctx.Value(2); v!= nil {
	fmt.Println(v)
	}else  {
		fmt.Println("no value associated with that key")
	}
}


