package blog

import "time"

var Version = "v1"

type Blog struct {
	Id         int64      `json:"id,string"`
	Uid        int64      `json:"uid,string"`
	Title      string     `json:"title"`
	Content    string     `json:"content"`
	CreateTime int64      `json:"create_time"`
	Ctime      *time.Time `json:"-"`
	Utime      *time.Time `json:"-"`
}

func (receiver *Blog) Version() string {
	return Version
}

func New() *Blog {
	return new(Blog)
}
