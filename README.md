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



