package entity

import "time"

type Blog struct {
	Id         int64      `json:"id,string"`
	Uid        int64      `json:"uid,string"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreateTime int64      `json:"create_time"`
	Ctime      *time.Time `json:"-"`
	Utime      *time.Time `json:"-"`
}

func (m *Blog) GetID() interface{} {
	return m.Id
}
