package harddrive


 import "https://github.com/monkrus/golang_architecture"


//DB 2
type Db map [int] architecture.User


// Save data function DB2
func (hd Db) Save (n int, u architecture.User) {
    m[n] = u
}
// retrievie data function DB2
func (hd Db)  Retrieve (n int) architecture.User {
    return m[n]
}