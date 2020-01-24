package architecture

// users
type User struct {
    First string
}



// specifies a set of methods
type  Accessor interface {
    Save (n int, u user)
    Retrieve (n int) user
}

//usage
func Put(a Accessor, n int, u user) {
    a.Save(n, u)
}

func Get(a Accessor, n int)  user {
    return a.retrieve(n)
}


