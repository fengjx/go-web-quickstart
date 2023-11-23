package entity

import (
	"time"
)

// User 用户信息表
// auto generate by gen cmd tool
type User struct {
	ID       int64     `json:"id"`       // -
	Username string    `json:"username"` // 用户名
	Pwd      string    `json:"pwd"`      // 密码
	Salt     string    `json:"salt"`     // 密码盐
	Nick     string    `json:"nick"`     // 昵称
	Utime    time.Time `json:"utime"`    // 更新时间
	Ctime    time.Time `json:"ctime"`    // 创建时间
}

func (m *User) GetID() interface{} {
	return m.ID
}
