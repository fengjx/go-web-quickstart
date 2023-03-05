package response

var (
	CodeOK        = 0
	CodeSystemErr = 10000
	CodeUserErr   = 20000
	CodeFacadeErr = 30000
)

// Resp
// Code 定义：错误产生来源+四位数字编号(业务定义)，返回0表示成功
// Code 1: 用户错误，2: 系统错误，3: 调用第三方系统错误
type Resp struct {
	Code int
	Msg  string
	Data interface{}
}

func OK() *Resp {
	return newBody(CodeOK, "ok")
}

func Data(data interface{}) *Resp {
	r := newBody(CodeOK, "ok")
	r.Data = data
	return r
}

func SysError(msg string) *Resp {
	return newBody(CodeSystemErr, msg)
}

func UserError(msg string) *Resp {
	return newBody(CodeUserErr, msg)
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
