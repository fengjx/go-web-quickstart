package utils

import "testing"

func TestToString(t *testing.T) {
	t.Log(ToString(1000))
	t.Log(ToString(1000.12))
	t.Log(ToString(map[string]string{"hello": "world"}))
}
