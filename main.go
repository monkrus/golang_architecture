package main


import "fmt"

//concrete type person
type person struct {
first string
}

func (p person) speak() {
fmt.Println ("My last warning, ", p.first)
}


type secretAgent struct {
person
ltk bool	
}

func (sa secretAgent) speak() {
fmt.Println ("Run now, " , sa.first)
}

// any TYPE that has the methods specified by an interface
// is also of the interface type
// an interface says
// "Hey, baby, if you have these methods, then you`re my type."

type human interface {
speak()
}

func main() {

//concrete type is person, secretAgent
//ABSTRACT type is human
p1 := person{ 
	first:"Miss Moneypenny",
}

sa1 := secretAgent {
	person: person {
		first:"James",	
	},
	ltk: true,
}
			

	fmt.Printf("%T\n", p1)


//1. Type person created
//2. Attached a method speak() with receiver (p person)
//3. Created interface human
//   VALUE can be of more than one TYPE
//   p1 is both TYPE person and TYPE human


var x, y human
x = p1
y = sa1
x.speak()
y.speak()
}
