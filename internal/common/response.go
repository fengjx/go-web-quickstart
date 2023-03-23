package common

import (
	"github.com/fengjx/go-web-quickstart/internal/common/env"
	"github.com/pkg/errors"
)

const (
	CodeOK   = 0
	CodeFail = 1
)

type RespCode int

var (
	CodeUserErr         RespCode = 20000
	CodeErrUnauthorized RespCode = 20001
	CodeErrBadRequest   RespCode = 20002
	CodeErrNotFound     RespCode = 20003
)

const (
	CodeSystemErr RespCode = 10000
)

const (
	CodeFacadeErr RespCode = 30000
)

// Resp
// Code 定义：错误产生来源+四位数字编号(业务定义)，返回0表示成功
// Code 1: 用户错误，2: 系统错误，3: 调用第三方系统错误
type Resp struct {
	Code RespCode    `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func Status(ok bool) *Resp {
	if ok {
		return OK()
	}
	return Fail()
}

func OK() *Resp {
	return newBody(CodeOK, "ok")
}

func Fail() *Resp {
	return newBody(CodeFail, "fail")
}

func Data(data interface{}) *Resp {
	r := newBody(CodeOK, "ok")
	r.Data = data
	return r
}

func Error(err error) *Resp {
	cause := errors.Cause(err)
	serviceErr, ok := cause.(ServiceError)
	if ok {
		return newBody(serviceErr.GetCode(), serviceErr.GetMsg())
	}
	msg := "server error"
	if !env.IsProd() {
		msg = err.Error()
	}
	return newBody(CodeSystemErr, msg)
}

func SysError(msg string) *Resp {
	return newBody(CodeSystemErr, msg)
}

func UserError(msg string) *Resp {
	return newBody(CodeUserErr, msg)
}

func ErrorBadRequest() *Resp {
	return newBody(CodeErrBadRequest, "BadRequest")
}

func ErrorUnauthorized() *Resp {
	return newBody(CodeErrUnauthorized, "Unauthorized")
}

func ErrorNotFound() *Resp {
	return newBody(CodeErrNotFound, "NotFound")
}

func FacadeError(msg string) *Resp {
	return newBody(CodeFacadeErr, msg)
}

func newBody(code RespCode, msg string) *Resp {
	r := new(Resp)
	r.Code = code
	r.Msg = msg
	return r
}
