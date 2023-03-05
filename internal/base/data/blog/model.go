package blog

import "time"

type Blog struct {
	Id       int64
	Uid      int64
	Title    string
	Content  string
	UpdateAt *time.Time
	CreateAt *time.Time
}

func New() *Blog {
	return new(Blog)
}
