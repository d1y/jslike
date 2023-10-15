package jsarray

import (
	"fmt"
	"testing"
)

func TestArray(t *testing.T) {
	var arr JSArray[int] = []int{1, 2, 3}
	val := arr.Splice(1, 1)
	flag, check := arr.At(-1)
	if val.Length() != 1 || !check || flag != 3 {
		t.FailNow()
	}

	var arr1 JSArray[int] = []int{}
	arr1.Push(12, 24, 32)
	arr1.Pop()
	arr1.Reverse()
	fmt.Println(arr1)
}
