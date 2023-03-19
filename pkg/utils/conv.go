package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func ToString(v interface{}) string {
	if v == nil {
		return ""
	}
	switch v.(type) {
	case string:
		return v.(string)
	case int, int8, int16, int32, int64:
		return fmt.Sprintf("%d", v)
	case float32, float64:
		return fmt.Sprintf("%f", v)
	case bool:
		if b, ok := v.(bool); ok && b {
			return "true"
		}
		return "false"
	default:
		return fmt.Sprintf("%#v", v)
	}
}

func ToInt(v interface{}) int {
	if v == nil {
		return 0
	}
	switch t := v.(type) {
	case float32:
		return int(t)
	case float64:
		return int(t)
	case int8:
		return int(t)
	case int16:
		return int(t)
	case int32:
		return int(t)
	case int64:
		return int(t)
	case uint8:
		return int(t)
	case uint16:
		return int(t)
	case uint32:
		return int(t)
	case uint64:
		return int(t)
	case int:
		return t
	case bool:
		if t {
			return 1
		}
		return 0
	case string:
		x := strings.TrimSpace(t)
		if x == "" {
			return 0
		}
		if strings.Contains(x, ".") {
			f, _ := strconv.ParseFloat(x, 64)
			return int(f)
		}
		res, _ := strconv.ParseInt(x, 10, 64)
		return int(res)
	}
	return 0
}

func ToInt64(v interface{}) int64 {
	if v == nil {
		return 0
	}
	switch t := v.(type) {
	case int:
		return int64(t)
	case int8:
		return int64(t)
	case int16:
		return int64(t)
	case int32:
		return int64(t)
	case int64:
		return t
	case uint:
		return int64(t)
	case uint8:
		return int64(t)
	case uint16:
		return int64(t)
	case uint32:
		return int64(t)
	case uint64:
		return int64(t)
	case float32:
		return int64(t)
	case float64:
		return int64(t)
	case bool:
		if t {
			return 1
		}
		return 0
	case []byte:
		return ToInt64(string(t))
	case string:
		t = strings.TrimSpace(t)
		if t == "" {
			return 0
		}
		if strings.Contains(t, ".") {
			f, _ := strconv.ParseFloat(t, 64)
			return int64(f)
		}
		res, _ := strconv.ParseInt(t, 10, 64)
		return res
	}
	return 0
}
