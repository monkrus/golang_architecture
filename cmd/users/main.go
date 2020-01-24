package main

import ( "fmt"
	"github.com/monkrus/golang_architecture"
	"github.com/monkrus/golang_architecture/storage/harddrive"
	
)

func main () {
	storage := harddrive{}

	u1 := architecture.User {
		first: "James",
	}

	u2 := architecture.User {
		first: "Jenny",
	}

	architecture.Put(storage, 1, u1)
	architecture.Put(storage, 2, u2)

	fmt.Println(architecture.Get(storage, 1))
	fmt.Println(architecture.Get(storage, 2))

}
