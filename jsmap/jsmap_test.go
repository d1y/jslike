package jsmap

import "testing"

func TestMap(t *testing.T) {
	var mp = New[string, int]()
	mp.Set("id", 12)
	if mp.Has("id") {
		id, _ := mp.Get("id")
		if id != 12 {
			t.FailNow()
		}
	}
}
