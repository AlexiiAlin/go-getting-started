package models

type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User // collection on users
	nextID = 1     // int32 data type
)
