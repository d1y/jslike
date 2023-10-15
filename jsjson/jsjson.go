package jsjson

import "encoding/json"

var (
	marshal   = json.Marshal
	unmarshal = json.Unmarshal
)

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/JSON/parse
func Parse[T any](data string) (*T, error) {
	var result = new(T)
	if err := unmarshal([]byte(data), result); err != nil {
		return nil, err
	}
	return result, nil
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify
func Stringify[T any](data T) (string, error) {
	val, err := marshal(data)
	if err != nil {
		return "", err
	}
	return string(val), nil
}
