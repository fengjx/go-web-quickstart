package common

import (
	"fmt"
)

type ServiceError interface {
	Error() string
	GetCode() RespCode
	GetMsg() string
}

type ServiceErr struct {
	Code RespCode
	Msg  string
}

func (receiver *ServiceErr) Error() string {
	return fmt.Sprintf("code: %d, msg: %s", receiver.GetCode(), receiver.GetMsg())
}

func (receiver *ServiceErr) GetCode() RespCode {
	return receiver.Code
}

func (receiver *ServiceErr) GetMsg() string {
	return receiver.Msg
}

func NewServiceErr(code RespCode, msg string) *ServiceErr {
	return &ServiceErr{
		Code: code,
		Msg:  msg,
	}
}
