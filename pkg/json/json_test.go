package json

import "testing"

func TestGetPathVal(t *testing.T) {
	jsonStr := `{"code":1,"message":"success","result":{"totalMsgCount":0},"time":"2023-03-18 23:30:54"}`
	t.Log(GetPathVal(jsonStr, "time").ToString())
}
