package jslike

import (
	"encoding/base64"
	"net/url"
	"strconv"
	"strings"
)

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/parseInt
// !
// ! not impl js parseInt function, only alias strconv.Atoi parse(10 base)
// !
func ParseInt(str string) (int, error) {
	return strconv.Atoi(str)
}

func MustParseInt(str string) int {
	v, _ := ParseInt(str)
	return v
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/parseFloat
// !
// ! not impl js parseInt function, only alias strconv.ParseFloat
// !
func ParseFloat(str string) (float64, error) {
	return strconv.ParseFloat(str, 64)
}

func MustParseFloat(str string) float64 {
	v, _ := ParseFloat(str)
	return v
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/encodeURIComponent
func EncodeURIComponent(str string) string {
	es := url.QueryEscape(str)
	es = strings.ReplaceAll(es, "+", "%20")
	es = strings.ReplaceAll(es, "%21", "!")
	es = strings.ReplaceAll(es, "%27", "'")
	es = strings.ReplaceAll(es, "%28", "(")
	es = strings.ReplaceAll(es, "%29", ")")
	es = strings.ReplaceAll(es, "%2A", "*")
	return es
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/decodeURIComponent
func DecodeURIComponent(raw string) (string, error) {
	t, err := url.QueryUnescape(strings.ReplaceAll(strings.ReplaceAll(string(raw), "+", "%2B"), "%20", "+"))
	if err != nil {
		return "", err
	}
	return t, nil
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/encodeURI
func EncodeURI(raw string) string {
	es := EncodeURIComponent(raw)
	ss := "!#$&'()*+,-./:=?@_~"
	for i := 0; i < len(ss); i++ {
		es = strings.ReplaceAll(es, "%"+strings.ToUpper(string(hexEncode([]byte(string(ss[i]))))), string(ss[i]))
	}
	return strings.ReplaceAll(es, "%3B", ";")
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/decodeURI
func DecodeURI(raw string) (string, error) {
	es := raw
	ss := "!#$&'()*+,-./:=?@_~"
	for i := 0; i < len(ss); i++ {
		es = strings.ReplaceAll(es, "%"+strings.ToUpper(string(hexEncode([]byte(string(ss[i]))))), "$"+"%"+strings.ToUpper(string(hexEncode([]byte(string(ss[i])))))+"$")
	}
	es, err := DecodeURIComponent(es)
	if err != nil {
		return "", err
	}
	for i := 0; i < len(ss); i++ {
		es = strings.ReplaceAll(es, "$"+string(ss[i])+"$", "%"+strings.ToUpper(string(hexEncode([]byte(string(ss[i]))))))
	}
	return es, nil
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/btoa
func Btoa(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}

// https://developer.mozilla.org/zh-CN/docs/Web/API/atob
func Atob(raw string) (string, error) {
	str, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return "", err
	}
	return string(str), nil
}
