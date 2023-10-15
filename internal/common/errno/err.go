package errno

import (
	"fmt"
)

type Errno struct {
	Code int
	Msg  string
}

func (e *Errno) Error() string {
	return fmt.Sprintf("code:%d, msg:%s", e.Code, e.Msg)
}

func (e *Errno) GetCode() int {
	return e.Code
}

func (e *Errno) GetMsg() string {
	return e.Msg
}

func NewErr(code int, msg string) *Errno {
	return &Errno{
		Code: code,
		Msg:  msg,
	}
}
