package main

import ("fmt")

// users
type user struct {
    first string
}
// DB 1, key (int), value user (string)
type mongo map [int] user

// save data function DB1
func (m mongo) save (n int, u user) {
    m[n] = u
}

// retrieve data function DB1
func (m mongo) retrieve (n int) user {
    return m[n]
}

//DB 2
type harddrive map [int] user


// save data function DB2
func (hd harddrive) save (n int, u user) {
    hd[n] = u
}
// retrievie data function DB2
func (hd harddrive)  retrieve (n int) user {
    return hd[n]
}
// specifies a set of methods
type  accessor interface {
    save (n int, u user)
    retrieve (n int) user
}

//usage
func put(a accessor, n int, u user) {
    a.save(n, u)
}

func get(a accessor, n int)  user {
    return a.retrieve(n)
}

func main () {
 storage := harddrive{}

u1 := user {
first: "James",
}

u2 := user {
    first: "Jenny",
}

put(storage, 1, u1)
put(storage, 2, u2)

fmt.Println(get(storage, 1))
fmt.Println(get(storage, 2))

}
