package entity

import (
	"time"
)

// Blog 博客表
// auto generate by gen cmd tool
type Blog struct {
	ID         int64     `json:"id"`          // -
	UID        int64     `json:"uid"`         // 用户ID
	Title      string    `json:"title"`       // 标题
	Content    string    `json:"content"`     // 内容
	CreateTime int64     `json:"create_time"` // 创建时间
	Utime      time.Time `json:"utime"`       // 更新时间
	Ctime      time.Time `json:"ctime"`       // 创建时间
}

func (m *Blog) GetID() interface{} {
	return m.ID
}
