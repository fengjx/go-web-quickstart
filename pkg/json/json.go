package json

import (
	jsoniter "github.com/json-iterator/go"
)

func ToJson(data interface{}) string {
	b, err := jsoniter.Marshal(data)
	if err != nil {
		return ""
	}
	return string(b)
}

func FromJson(jsonStr string, target interface{}) error {
	err := jsoniter.Unmarshal([]byte(jsonStr), &target)
	if err != nil {
		return err
	}
	return nil
}

func GetPathVal(jsonStr string, path interface{}) any {
	return jsoniter.Get([]byte(jsonStr), path)
}
