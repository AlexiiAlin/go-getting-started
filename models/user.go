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

func GetUsers() []*User {
	return users
}

func AddUser(u User) (User, error) {
	u.ID = nextID
	nextID++
	users = append(user, &u)
	return u, nil
}
