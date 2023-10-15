package response

import (
	"errors"

	"github.com/fengjx/go-web-quickstart/internal/common/env"
	"github.com/fengjx/go-web-quickstart/internal/common/errno"
)

const (
	CodeOK = 0
)

var (
	// CodeUserErr 通用用户错误
	CodeUserErr         = 20000
	CodeErrUnauthorized = 20401
	CodeErrBadRequest   = 20400
	CodeErrNotFound     = 20404
)

const (
	// CodeSystemErr 通用系统内部错误
	CodeSystemErr = 10000
)

const (
	// CodeFacadeErr 通用系统内部错误
	CodeFacadeErr = 30000
)

// Resp
// Code 定义：错误产生来源+四位数字编号(业务定义)，返回0表示成功
// Code 1: 用户错误，2: 系统错误，3: 调用第三方系统错误
type Resp struct {
	Code int         `json:"code"`
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
	return newBody(CodeSystemErr, "fail")
}

func Data(data interface{}) *Resp {
	r := newBody(CodeOK, "ok")
	r.Data = data
	return r
}

func Error(err error) *Resp {
	var errn *errno.Errno
	if ok := errors.As(err, &errn); ok {
		return newBody(errn.GetCode(), errn.GetMsg())
	}
	msg := "system error"
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

func ErrorBadRequestWithMsg(msg string) *Resp {
	return newBody(CodeErrBadRequest, msg)
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

func newBody(code int, msg string) *Resp {
	r := new(Resp)
	r.Code = code
	r.Msg = msg
	return r
}
