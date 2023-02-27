package user

type User struct {
}

func New() *User {
	return new(User)
}
