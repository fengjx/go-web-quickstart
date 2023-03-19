package user

import "time"

var Version = "v1"

type User struct {
	Id       int64      `json:",string"`
	Username string     `json:"username"`
	Pwd      string     `json:"-"`
	Salt     string     `json:"-"`
	Nick     string     `json:"nick"`
	Ctime    *time.Time `json:"-"`
	Utime    *time.Time `json:"-"`
}

func (receiver *User) Version() string {
	return Version
}

func New() *User {
	return new(User)
}
