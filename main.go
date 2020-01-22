package main


import "fmt"

type person struct {

first string
}



func (p person) speak() {

fmt.Println ("from a person - this is my name", p.first)
}


// any TYPE that has the methods specified by an interface
// is also of the interface type
// an interface says
// "Hey, baby, if you have these methods, then you`re my type."


type human interface {
	speak()
}


func main() {

		p1 := person{
			first:"James",
			
		}

	fmt.Printf("%T\n", p1)
}

//1. Type person created
//2. Attached a method speak() with receiver (p person)
//3. Created interface human
//   VALUE can be of more than one TYPE
//   p1 is both TYPE person and TYPE human


var x human
x = p1
fmt.Printf("%T\n", x)