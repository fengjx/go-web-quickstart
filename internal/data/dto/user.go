package dto

import (
	"time"

	"github.com/fengjx/go-web-quickstart/internal/data/entity"
)

type UserDTO struct {
	ID       int64     `json:"id,string"`
	Username string    `json:"username"`
	Nick     string    `json:"nick"`
	Ctime    time.Time `json:"ctime"`
}

func (dto *UserDTO) Of(model *entity.User) {
	dto.ID = model.ID
	dto.Username = model.Username
	dto.Nick = model.Nick
	dto.Ctime = model.Ctime
}
