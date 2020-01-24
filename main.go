package main

import "fmt"

//1.create  type HOTDOG
type hotdog int

//2. create a function
// this method `implements` the type HOTDOG
func(h hotdog) cook() {
    fmt.Println("I am cooking")
}

//3. create type HOTFOOD interface
// Note it is also `using` a cook() method 
type hotFood interface {
cook()
}


func main() {
var x hotdog = 42
fmt.Printf ("%T\n", x)
var y hotFood
fmt.Printf ("%T\n", y)
y = x
fmt.Printf ("%T\n", y)
}