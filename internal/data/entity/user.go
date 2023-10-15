package entity

import "time"

type User struct {
	Id       int64      `json:"id,string"`
	Username string     `json:"username"`
	Pwd      string     `json:"pwd"`
	Salt     string     `json:"salt"`
	Nick     string     `json:"nick"`
	Ctime    *time.Time `json:"ctime"`
	Utime    *time.Time `json:"utime"`
}

func (m *User) GetID() interface{} {
	return m.Id
}
