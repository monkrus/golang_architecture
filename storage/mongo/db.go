package mongo


 import "https://github.com/monkrus/golang_architecture"
// DB 1, key (int), value User (string)
type Db map [int] User

// Save data function DB1
func (m Db) Save (n int, u User) {
    m[n] = u
}

// retrieve data function DB1
func (m Db) Retrieve (n int) User {
    return m[n]
}