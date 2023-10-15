package array

import (
	"fmt"
	"strings"
)

// copy by https://github.com/prashantacharya/js-like-functions/blob/main/arrays/arrays.go

type JSArray[T comparable] []T

// ===========not standard function

func (v JSArray[T]) IsEmpty() bool {
	return len(v) == 0
}

func (v JSArray[T]) IsNotEmpty() bool {
	return !v.IsEmpty()
}

// ======================================

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/length
func (v JSArray[T]) Length() int {
	return len(v)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/forEach
func (v JSArray[T]) Foreach(f func(T)) {
	for _, val := range v {
		f(val)
	}
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/map
func (v JSArray[T]) Map(f func(T) T) JSArray[T] {
	var mappedValues JSArray[T]
	for _, val := range v {
		mappedValues = append(mappedValues, f(val))
	}
	return mappedValues
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/filter
func (v JSArray[T]) Filter(f func(T) bool) JSArray[T] {
	var filteredValues JSArray[T]
	for _, val := range v {
		if f(val) {
			filteredValues = append(filteredValues, val)
		}
	}

	return filteredValues
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/includes
func (v JSArray[T]) Includes(val T) bool {
	for _, value := range v {
		if value == val {
			return true
		}
	}

	return false
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/at
func (v JSArray[T]) At(n int) (T, bool) {
	if n < 0 {
		n += v.Length()
	}
	if n < 0 || n >= v.Length() {
		return *new(T), false
	}
	if v.Length() == 1 {
		return *new(T), true
	}
	return v[n], true
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/splice
func (v *JSArray[T]) Splice(start int, delete int, item ...T) JSArray[T] {
	deleteVal := splice[T](v, start, delete, item)
	return deleteVal
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/push
func (v *JSArray[T]) Push(item ...T) {
	*v = append(*v, item...)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/unshift
func (v *JSArray[T]) Unshift(item ...T) {
	*v = append(item, *v...)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/pop
func (v *JSArray[T]) Pop() (T, bool) {
	if v.Length() == 0 {
		return *new(T), false
	}
	return v.Splice(v.Length()-1, 1)[0], true
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/shift
func (v *JSArray[T]) Shift() (T, bool) {
	if v.Length() == 0 {
		return *new(T), false
	}
	return v.Splice(0, 1)[0], true
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/reverse
func (v *JSArray[T]) Reverse() {
	for i, j := 0, v.Length()-1; i < j; i, j = i+1, j-1 {
		(*v)[i], (*v)[j] = (*v)[j], (*v)[i]
	}
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/toReversed
func (v *JSArray[T]) ToReversed() *JSArray[T] {
	var temp = make(JSArray[T], v.Length())
	copy(*v, temp)
	temp.Reverse()
	return &temp
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/join
func (v JSArray[T]) Join(sep string) string {
	var strArr = make([]string, len(v))
	for _, item := range v {
		strArr = append(strArr, fmt.Sprintf("%v", item))
	}
	return strings.Join(strArr, sep)
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/toString
func (v JSArray[T]) ToString() string {
	return v.Join(",")
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/with
func (v *JSArray[T]) With(index int, value T) *JSArray[T] {
	var dist = make(JSArray[T], v.Length())
	copy(dist, *v)
	dist[index] = value
	return &dist
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/some
func (v *JSArray[T]) Some(f func(T) bool) bool {
	for _, val := range *v {
		if f(val) {
			return true
		}
	}
	return false
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/every
func (v *JSArray[T]) Every(f func(T) bool) bool {
	for _, val := range *v {
		if !f(val) {
			return false
		}
	}
	return true
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/indexOf
func (v *JSArray[T]) IndexOf(val T) int {
	for i, item := range *v {
		if item == val {
			return i
		}
	}
	return -1
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/lastIndexOf
func (v *JSArray[T]) LastIndexOf(val T) int {
	for i := len(*v) - 1; i >= 0; i-- {
		if (*v)[i] == val {
			return i
		}
	}
	return -1
}

// https://developer.mozilla.org/zh-CN/docs/Web/JavaScript/Reference/Global_Objects/Array/concat
func (v *JSArray[T]) Concat(appendValue JSArray[T]) JSArray[T] {
	return append(*v, appendValue...)
}
