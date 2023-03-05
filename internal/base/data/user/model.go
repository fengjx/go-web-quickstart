package user

import "time"

var Version = "v1"

type User struct {
	Id       int64 `json:",string"`
	Username string
	Pwd      string
	Salt     string
	Nick     string
	Ctime    *time.Time
	Utime    *time.Time
}

func New() *User {
	return new(User)
}
