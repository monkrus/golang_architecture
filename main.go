package main

import (
	"fmt"
	"log"
)

//Type error implements the ERROR interface
//implement the string
// takes receiver`string`, implements `Error()`, returns `string`
// we are getting an error message, NOT creating a new error

type errorString string

func (es errorString) Error() string {

	return fmt.Sprintf("this is an error string with info -%s", string(es))
}

func main() {
	n, err := sum(2, 4)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(n)
}

func sum(i, j int) (int, error) {
	n := i + j
	if n != i+j {
		var sErr errorString = "the numbers did not add up"
		return 0, sErr
	}
	return n, nil
}

//error interface
