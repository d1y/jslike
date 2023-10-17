package jsstring

import (
	"math"
	"strings"
	"unicode"
)

type JSString string

// ===========not standard function
func (s JSString) IsEmpty() bool {
	return len(s) == 0
}

func (s JSString) NotEmpty() bool {
	return len(s) >= 1
}

func (s JSString) Equal(other string) bool {
	return string(s) == other
}

func (s JSString) NotEqual(other string) bool {
	return !s.Equal(other)
}

// https://github.com/bytedance/gopkg/blob/3db87571198b1ff8824d3ed695e3167e4cb3d699/lang/stringx/is.go#L48C1-L58C2
func (s JSString) IsNumeric() bool {
	if s == "" {
		return false
	}
	for _, v := range s {
		if !unicode.IsDigit(v) {
			return false
		}
	}
	return true
}

func (s JSString) ToNative() string {
	return string(s)
}

// ======================================

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/toLowerCase
func (s JSString) ToLowerCase() JSString {
	return JSString(strings.ToLower(string(s)))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/toUpperCase
func (s JSString) ToUpperCase() JSString {
	return JSString(strings.ToUpper(string(s)))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/split
// ** unfull impl this function
func (s JSString) Split(sep string) []JSString {
	arr := strings.Split(string(s), sep)
	result := make([]JSString, len(arr))
	for idx, item := range arr {
		result[idx] = JSString(item)
	}
	return result
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/length
func (s JSString) Length() int {
	return len(s)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/concat
func (s JSString) Concat(strs ...string) JSString {
	result := string(s)
	for _, str := range strs {
		result += str
	}
	return JSString(result)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/endsWith
func (s JSString) EndsWith(suffix string, pos ...int) bool {
	position := 0
	if len(pos) == 1 {
		position = pos[0]
	}
	if position >= s.Length() {
		return false
	}
	return strings.HasSuffix(string(s)[position:], suffix)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/startsWith
func (s JSString) StartsWith(prefix string, pos ...int) bool {
	position := 0
	if len(pos) == 1 {
		position = pos[0]
	}
	if position >= s.Length() {
		return false
	}
	return strings.HasPrefix(string(s)[position:], prefix)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/includes
func (s JSString) Includes(str string, pos ...int) bool {
	position := -1
	if len(pos) == 1 {
		position = pos[0]
	}
	if position >= s.Length() {
		return false
	}
	if position >= 0 {
		return strings.Contains(string(s)[position:], str)
	}
	return strings.Contains(string(s), str)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/at
// impl by https://github.com/tc39/proposal-relative-indexing-method#polyfill
//
// refer to: https://github.com/mathiasbynens/String.prototype.at
func (s JSString) At(n int) JSString {
	if n < 0 {
		n += s.Length()
	}
	if n < 0 || n >= s.Length() {
		return JSString("")
	}
	return JSString(string(s)[n : n+1])
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/indexOf
func (s JSString) IndexOf(searchValue string, pos ...int) int {
	position := 0
	if len(pos) == 1 {
		position = pos[0]
	}
	if position >= s.Length() {
		return -1
	}
	return strings.Index(string(s)[position:], searchValue)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/lastIndexOf
func (s JSString) LastIndexOf(searchValue string, pos ...int) int {
	position := 0
	if len(pos) == 1 {
		position = pos[0]
	}
	if position >= s.Length() {
		return -1
	}
	return strings.LastIndex(string(s)[position:], searchValue)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padEnd
// https://github.com/tc39/proposal-string-pad-start-end/blob/main/polyfill.js
func (s JSString) PadEnd(targetLength int, padString ...string) JSString {
	if s.Length() >= targetLength {
		return s
	}
	var fillString = ""
	if len(padString) == 1 {
		fillString = padString[0]
	}
	fillLen := targetLength - len(s)
	fillCount := fillLen / len(fillString)
	truncatedFillString := fillString[:fillLen%len(fillString)]
	fill := strings.Repeat(fillString, fillCount) + truncatedFillString
	return s.Concat(fill)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/padStart
// https://github.com/tc39/proposal-string-pad-start-end/blob/main/polyfill.js
func (s JSString) PadStart(targetLength int, padString ...string) JSString {
	if s.Length() >= targetLength {
		return s
	}
	var fillString = ""
	if len(padString) == 1 {
		fillString = padString[0]
	}
	fillLen := targetLength - len(s)
	fillCount := fillLen / len(fillString)
	truncatedFillString := fillString[:fillLen%len(fillString)]
	fill := strings.Repeat(fillString, fillCount) + truncatedFillString
	return JSString(fill + string(s))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/repeat
func (s JSString) Repeat(count int) JSString {
	return JSString(strings.Repeat(string(s), count))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/replace
// !
// ! this is sample impl
// ! only support string replace
// !
func (s JSString) Replace(pattern string, replacement string) JSString {
	return JSString(strings.Replace(string(s), pattern, replacement, -1))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/replaceAll
// !
// ! this is sample impl
// ! only support string replaceAll
// !
func (s JSString) ReplaceAll(pattern string, replacement string) JSString {
	return JSString(strings.ReplaceAll(string(s), pattern, replacement))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/slice
func (s JSString) Slice(indexStart int, indexEnds ...int) JSString {
	if indexStart > s.Length() {
		return ""
	}
	if indexStart < 0 {
		indexStart = int(math.Max(float64(indexStart+s.Length()), 0))
	}
	indexEnd := len(s)
	if len(indexEnds) == 1 {
		indexEnd = indexEnds[0]
	}
	if indexEnd <= 0 {
		indexEnd = int(math.Max(float64(indexEnd+s.Length()), 0))
	}
	if indexEnd > s.Length() {
		indexEnd = s.Length() // 这里的行为应该跟js不一致 !_!
	}
	if indexEnd <= indexStart {
		return ""
	}
	return JSString(string(s)[indexStart:indexEnd])
}

func (s JSString) Substring(indexStart int, indexEnds ...int) JSString {
	if indexStart > s.Length() {
		return ""
	}
	if len(indexEnds) == 1 {
		indexEnd := indexEnds[0]
		if indexEnd == 0 {
			return ""
		}
		if indexEnd > s.Length() {
			indexEnd = s.Length()
		}
		return s[indexStart:indexEnd]
	}
	return s[indexStart:]
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/trim
func (s JSString) Trim() JSString {
	return JSString(strings.TrimSpace(string(s)))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/trimStart
func (s JSString) TrimStart() JSString {
	return JSString(strings.TrimLeft(string(s), " "))
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/String/trimEnd
func (s JSString) TrimEnd() JSString {
	return JSString(strings.TrimRight(string(s), " "))
}
