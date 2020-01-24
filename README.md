### GO

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









