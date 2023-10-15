package jsset

import (
	"testing"
)

func TestJsNew(t *testing.T) {
	ms := New[int]()
	ms.Add(12)
	ms.Add(42)
	ms.Add(120)
	has1 := ms.Has(120)
	if !has1 {
		t.FailNow()
	}
	if ms.Size() != 3 {
		t.FailNow()
	}
}
