package main

import (
    "fmt"
)


type person struct {
	first string
}

type mongo   map [int]person
type postg   map [int]person

// `m mongo` is a receiver
// `n int, p person` is a key/value pair
func (m mongo) save (n int, p person) {
	m[n] = p
}

func (m mongo) retreive (n int) person {
	return m[n]
}

func (pg postg) save (n int, p person) {
	pg[n] = p
}

func (pg postg) retreive (n int)  person {
	return pg[n]
}

type accessor interface {
	save(n int, p person)
	retreive(n int) person
}

func put (a accessor, n int, p person) {
	a.save(n, p)
}

func get(a accessor, n int) person {
return	a.retreive(n)
}

func main() {

	dbm := mongo{}
    dbp := postg{}

    p1:= person {
    	first: "Jenny",
    }

    p2:= person {
    	first: "James",
    }

    put(dbm, 1, p1)
    put(dbm, 2, p2)

    fmt.Println(get(dbm, 1))
    fmt.Println(get(dbm, 2))

    //or store in some other data storage
    put(dbp, 1, p1)
    put(dbp, 2, p2)

    fmt.Println(get(dbp, 1))
    fmt.Println(get(dbp, 2))

    }
