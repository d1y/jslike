package jspromise

import (
	"testing"
	"time"
)

func TestPromise(t *testing.T) {
	promise := New[int](func() (int, error) {
		time.Sleep(time.Second * 2)
		return 12, nil
	})
	val, err := promise.Await()
	if err != nil || val != 12 {
		t.FailNow()
	}
	e, _ := promise.Read()
	if e != 12 {
		t.FailNow()
	}
}
