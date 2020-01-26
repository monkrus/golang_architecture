package main

import (
"github.com/monkrus/golang_architecture/session"
)

func main() {
	ctx:= context.Background()
	ctx = session.SetUserId(ctx, 1)
	ctx = session.SetIsAdmin(ctx, true)
	i:= session.GetUserId(ctx)
	b:= session.GetIsAdmin(ctx)
	fmt.Printf("user %d is admin %t\n", i, b)



}