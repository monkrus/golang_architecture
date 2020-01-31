idy### GO

## installation

1. (https://golang.org/dl/)

2. `go mod init`  creates a new go.mod file and automatically imports dependencies from Godeps.json, Gopkg.lock, or a number of other supported formats. (e.g. go mod init github.com/monkrus/golang_architecture.git)

3. `go mod tidy` finds all the packages transitively imported by packages in your module. It adds new module requirements for packages not provided by any known module, and it removes requirements on modules that don't provide any imported packages 


## tags

1. To add a tag, type `git tag  v0.0.1`, then `git log` and `git push --tags`

2. To get to any specific version, type `git pull`, then `git tag` (to see all the tags),
   'git checkout v0.0.2' and `git checkout master` to return to the normal state.

##curl

1. To see headers for the specific website, type `curl -v  https://www.google.com/teapot `
   (GET...Accept), (HTTP...Accept)

## web architecture

1. Organizing code: 

- visiblity (e.g. `type Person struct` starting with capital letter is visible from outiside of the package  )
- separation (create separate packages instead of one huge package)

- import packages `go get golang.org/x/tools/cmd/goimports`

- `cmd` folder cointains all the packages (e.g. package person) , and package person will contain `main.go` inside of it

- package documentation goes into `doc.go` file, types,function ,interface etc. documentation is added in comments above the code.

## notes

1. In Go, `:=`  is for declaration + assignment, whereas `=` is for assignment only.For example, var foo int = 10 is the same as foo := 10

2. Background returns a non-nil, empty Context. It is never canceled, has no values, and has no deadline. 
   It is typically used by the main function, initialization, and tests, and as the top-level Context for incoming requests.

3. functions : 
`func function_name( [parameter list] ) [return_types]
{
   body of the function
}`

-parameter is a placeholder for values (arguments) to pass.The parameter list refers to the type, order amd number of parameters in function.

-return type is a list of data type of the values the function returns.

4. `&` returns the memory address of the following variable.

   `*` returns the value of the following variable

## errors

-deal with your errors OR provide meaning so that your error tells what, where, how, whom, and when (as applicable) and return it

-if we just fmt.Println or log.Println or panic our “error” then we don’t get this extra meaning. 

ADD MEANING TO YOUR ERRORS WITH

-fmt.Errorf and use %w and pass in the error to %w
-this allows us to unwrap our error to one level of func calls
-this allows us to use IS or AS to many levels of func calls

-errors.New() when you don’t need to pass in a previous error
-good for creating a “base” error that everything is going off
-you can also use fmt.Errorf for the “base” error and just not pass into %w

-pass errors up through the call stack as necessary, adding something to each pass







