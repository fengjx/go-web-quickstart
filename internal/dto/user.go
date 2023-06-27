package dto

import (
	"github.com/fengjx/go-web-quickstart/internal/base/dao/user"
	"time"
)

type UserDTO struct {
	Id       int64      `json:"id,string"`
	Username string     `json:"username"`
	Nick     string     `json:"nick"`
	Ctime    *time.Time `json:"ctime"`
}

func (dto *UserDTO) Of(model *user.User) {
	dto.Id = model.Id
	dto.Username = model.Username
	dto.Nick = model.Nick
	dto.Ctime = model.Ctime
}
